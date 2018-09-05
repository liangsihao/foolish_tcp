package main

import (
	"conn"
	"flag"
)

func main() {
	flag.Parse()
	conn.RunTcpServer(&Myapp{})
}

type Myapp struct{}

func (this *Myapp) Request(ph *conn.PackHead) (*conn.PackHead, error) {
	return ph, nil
}

func (this *Myapp) Response(ph *conn.PackHead) error {
	return nil
}

func (this *Myapp) OnClose(){

}
