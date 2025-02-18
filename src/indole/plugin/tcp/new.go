package tcp

import (
	"io"
	"log"
	"net"
)

// New ...
func New(args *Args) io.ReadWriteCloser {
	conn, err := net.Dial(args.Network, args.Address)
	if err != nil {
		log.Println("[plugin]", "[tcp]", "[New]", "err:", err)
		return nil
	}
	return &TCP{
		conn: conn,
	}
}

// Args ...
type Args struct {
	Network string `xml:"network,attr"`
	Address string `xml:"address,attr"`
}

// NewByConn ...
func NewByConn(conn net.Conn) io.ReadWriteCloser {
	return &TCP{
		conn: conn,
	}
}
