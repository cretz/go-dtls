package dtls

import (
	"net"
	"time"
)

type Conn struct {
	conn   net.PacketConn
	config Config
	client bool
}

var _ net.PacketConn = &Conn{}

type Config struct {
}

func Client(conn net.PacketConn, config *Config) *Conn {
	c := &Conn{conn: conn, client: true}
	if config != nil {
		c.config = *config
	}
	return c
}

// Dial is the equivalent of calling Client followed by Handshake. Failures
// do not close the connection.
func Dial(conn net.PacketConn, config *Config, network, addr string) (*Conn, error) {
	c := Client(conn, config)
	if err := c.Handshake(network, addr); err != nil {
		return nil, err
	}
	return c, nil
}

func Server(conn net.PacketConn, config *Config) *Conn {
	c := &Conn{conn: conn, client: false}
	if config != nil {
		c.config = *config
	}
	return c
}

func (c *Conn) Handshake(network, addr string) error {
	// TODO: this is a handshake only if not already done
	panic("TODO")
}

func (c *Conn) IsEstablished(network, addr string) bool {
	panic("TODO")
}

func (c *Conn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	// TODO: this is inefficient, do better w/ a buffer or something
	var buf [508]byte
	n, addr, err = c.conn.ReadFrom(buf[:])
	if err != nil {
		return 0, addr, err
	}
	if c.client {
		n, err = c.clientReadFrom(buf[:n], addr, p)
	} else {
		n, err = c.serverReadFrom(buf[:n], addr, p)
	}
	return
}

func (c *Conn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	// TODO: assert connected via handshake, rehandshake if needed
	panic("TODO")
}

// Close calls the original PacketConn's Close. If the PacketConn is being
// reused by multiple Conns, users may not want to call this. There is no
// other cleanup that this does.
func (c *Conn) Close() error { return c.conn.Close() }

func (c *Conn) LocalAddr() net.Addr                { return c.conn.LocalAddr() }
func (c *Conn) SetDeadline(t time.Time) error      { return c.conn.SetDeadline(t) }
func (c *Conn) SetReadDeadline(t time.Time) error  { return c.conn.SetReadDeadline(t) }
func (c *Conn) SetWriteDeadline(t time.Time) error { return c.conn.SetWriteDeadline(t) }
