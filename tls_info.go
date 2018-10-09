package dtls

import (
	"github.com/cretz/go-dtls/model"
)

type tlsInfo interface {
	// Record can be nil to start client handshake
	handshake(rec *model.DTLSRecord) error
}

func newTLSInfo(conn *Conn, network, addr string) tlsInfo {
	if conn.client {
		return &clientTLSInfo{conn, network, addr}
	}
	panic("TODO: server")
}

type clientTLSInfo struct {
	conn    *Conn
	network string
	addr    string
}

func (c *clientTLSInfo) handshake(rec *model.DTLSRecord) error {
	// If the rec is nil, this wants to start from scratch
	if rec == nil {
		// Only start the handshake if we're not established
		if c.established() {
			return nil
		}
		return c.sendClientHello()
	}
}

func (c *clientTLSInfo) established() bool {
	panic("TODO")
}
