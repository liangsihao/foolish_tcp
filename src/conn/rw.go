package conn

import (
	"github.com/golang/glog"
	"net"
	"time"
)

var READ_BUFFER_SIZE = 64

type connection struct {
	conn   net.Conn
	client *TcpClient
}

// 愚蠢的读
func (this *connection) Reader(app App) {
	uid := uint32(0)
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln("recoving", err)
		}
		this.conn.Close()
		if uid != 0 {
			ClientMap.RemoveClient(uid)
		}
		app.OnClose()
	}()

	readLen := 0
	buf := make([]byte, READ_BUFFER_SIZE)
	for {
		// 愚蠢的超时
		this.conn.SetReadDeadline(time.Now().Add(READ_TIMEOUT))
		n, err := this.conn.Read(buf[readLen:])
		if err != nil {
			break
		}
		if n == 0 {
			break
		}

		// 愚蠢的伸缩，到达一半，增加一倍
		for (readLen + n) > len(buf)/2 {
			tmp := make([]byte, len(buf)*2)
			copy(tmp, buf)
			buf = tmp
		}
		for (readLen+n) > READ_BUFFER_SIZE && (readLen+n) < len(buf)/4 {
			tmp := make([]byte, len(buf)/2)
			copy(tmp, buf)
			buf = tmp
		}

		readLen += n
		ph, errorCode, err, le := UnmarshallHead(buf[:readLen], n)

		if errorCode == Marshall_Len_Less_fail {
			glog.Error("unmarshalling fail:", errorCode, err)
			continue
		}

		// 如果包头错误，读多少扔多少
		if errorCode == Marshall_Flag_fail {
			glog.Error("unmarshalling fail:", errorCode, err)
			buf = append(buf[readLen:], make([]byte, readLen)...)
			break
		}
		if errorCode != Marshall_Success {
			glog.Error("unmarshalling fail:", errorCode, err)
			buf = append(buf[n:], make([]byte, n)...)
			break
		}

		uid = ph.UserId
		ClientMap.NewClient(this.client, uid)
		ret, _ := app.Request(ph)
		this.Writer(ret, uid)

		// 愚蠢的清理buf
		buf = append(buf[le:], make([]byte, le)...)
		readLen -= le
	}
}

// 不讲究的写
func (this *connection) Writer(ph *PackHead, uid uint32) {
	n, err := this.conn.Write(MarshallHead(ph))
	if err != nil {
		glog.Infoln(n, err)
	}
}
