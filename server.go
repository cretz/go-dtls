package dtls

import "net"

type serverTLSInfo struct {
}

// If result is 0, nil then callers should read more immediately
func (c *Conn) serverReadFrom(b []byte, addr net.Addr, p []byte) (int, error) {
	panic("TODO")
}
