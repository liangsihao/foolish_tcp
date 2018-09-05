package conn

import (
	"net"
	"sync"
)

var ClientMap clientMap

type clientMap struct {
	l sync.RWMutex
	m map[uint32]*TcpClient
}

type TcpClient struct {
	conn   *connection
	userid uint32
}

func (this *TcpClient) NewTcpClient(c net.Conn, client *TcpClient) {
	this.conn = &connection{
		conn:   c,
		client: client,
	}
}

func init() {
	ClientMap.m = make(map[uint32]*TcpClient)
}

func (this *clientMap) NewClient(c *TcpClient, uid uint32) {
	this.l.Lock()
	defer this.l.Unlock()

	c.userid = uid
	this.m[uid] = c
}

func (this *clientMap) GetClient(uid uint32) *TcpClient {
	this.l.Lock()
	defer this.l.Unlock()

	if c, ok := ClientMap.m[uid]; ok {
		return c
	}
	return nil
}

func (this *clientMap) RemoveClient(uid uint32) {
	if _, ok := ClientMap.m[uid]; ok {
		this.l.Lock()
		delete(this.m, uid)
		this.l.Unlock()
	}
}
