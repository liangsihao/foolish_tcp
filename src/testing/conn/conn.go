package main

import (
	"bytes"
	"conn"
	"flag"
	"github.com/golang/glog"
	"net"
	"time"
)

var (
	NET_WORK     = "tcp"
	DIAL_ADDRESS = "localhost:9999"
	DIAL_TIMEOUT = 5 * time.Second
)

func main() {
	flag.Parse()

	cli := TcpClient(uint32(10086))
	cli.SendRecving("hey sihao")
}

func TcpClient(uid uint32) *conn.TcpClient {
	c, err := net.DialTimeout(NET_WORK, DIAL_ADDRESS, DIAL_TIMEOUT)
	if err != nil {
		glog.Errorln("connecting fail", err)
		return nil
	}

	cli := &conn.TcpClient{}
	cli.NewTcpClient(c, cli)
	return cli
}

func SendRecving(cli *conn.TcpClient,msg string) {
	body := bytes.NewBufferString(msg).Bytes()

	go func() {

		conn.Reader(this.c)
	}()

	for i := 0; i < 3; i++ {
		//	glog.Infoln(i)
		buf := conn.MarshallHead(&conn.PackHead{
			ParseFlag:  conn.FLAG,
			Length:     uint32(len(body)),
			RealLength: uint32(len(body)),
			UserId:     this.Uid,
			Code:       uint32(1),
			Body:       body,
		})
		this.c.Write(buf)
	}
	glog.Infoln("xxxs")
	//}
	time.Sleep(10 * time.Second)
	this.c.Close()
}
