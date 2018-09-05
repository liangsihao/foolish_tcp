package conn

import (
	"encoding/binary"
	"github.com/golang/glog"
)

var (
	FLAG     uint32 = 0x9263
	HEAD_LEN        = 32
)

type PackHead struct {
	ParseFlag  uint32 // [0:4]   package head
	Length     uint32 // [4:8]  length of the package
	RealLength uint32 // [8:12] real length for encode
	UserId     uint32 // [12:16] user id [16:24]
	ServerId   uint32 // [16:20] sid
	Tmp        uint32 // [20:24] sid
	Tmp2       uint32 // [24:28] sid
	Code       uint32 // [28:32] proto code
	Body       []byte // [32:]   body
}

func UnmarshallHead(buf []byte, readLen int) (*PackHead, int, error, int) {

	flag := binary.BigEndian.Uint32(buf[0:4])
	length := binary.BigEndian.Uint32(buf[4:8])
	bufLength := int(length) + HEAD_LEN
	if flag != FLAG {
		glog.Errorln(flag, FLAG)
		return nil, Marshall_Flag_fail, Marshall_Flag_Err, bufLength
	}
	if length > uint32(readLen-HEAD_LEN) {
		return nil, Marshall_Len_Less_fail, Marshall_Len_Less_Err, bufLength
	}
	realLength := binary.BigEndian.Uint32(buf[8:12])
	uid := binary.BigEndian.Uint32(buf[12:16])
	sid := binary.BigEndian.Uint32(buf[16:20])
	code := binary.BigEndian.Uint32(buf[28:32])

	return &PackHead{
		ParseFlag:  flag,
		Length:     length,
		RealLength: realLength,
		UserId:     uid,
		ServerId:   sid,
		Code:       code,
		Body:       buf[HEAD_LEN : HEAD_LEN+int(length)],
	}, Marshall_Success, nil, bufLength
}

func MarshallHead(head *PackHead) []byte {
	buf := make([]byte, int(head.Length)+HEAD_LEN)
	binary.BigEndian.PutUint32(buf[0:4], head.ParseFlag)
	binary.BigEndian.PutUint32(buf[4:8], head.Length)
	binary.BigEndian.PutUint32(buf[8:12], head.RealLength)
	binary.BigEndian.PutUint32(buf[12:16], head.UserId)
	binary.BigEndian.PutUint32(buf[16:20], head.ServerId)
	binary.BigEndian.PutUint32(buf[20:24], head.Tmp)
	binary.BigEndian.PutUint32(buf[24:28], head.Tmp2)
	binary.BigEndian.PutUint32(buf[28:32], head.Code)
	copy(buf[HEAD_LEN:], head.Body)

	return buf
}
