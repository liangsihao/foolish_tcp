package conn

import (
	"github.com/golang/glog"
	"net"
	"time"
)

var (
	NET_WORK       = "tcp"
	LISTEN_ADDRESS = "localhost:9999"
	READ_TIMEOUT   = 5 * time.Second
)

func RunTcpServer(app App) {
	listener, err := net.Listen(NET_WORK, LISTEN_ADDRESS)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			glog.Errorln("<error>", err)
			connection.Close()
			continue
		}
		go func() {
			client := &TcpClient{}
			client.NewTcpClient(connection,client)
			client.conn.Reader(app)
		}()
	}
}
