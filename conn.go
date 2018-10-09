package dtls

import (
	"net"
	"sync"
	"time"
)

type Conn struct {
	conn   net.PacketConn
	config Config
	client bool

	tlsInfoLock sync.Mutex
	tlsInfos    map[string]tlsInfo
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

// Handshake performs handshake only if not already done
func (c *Conn) Handshake(network, addr string) error {
	info, _ := c.tlsInfo(network, addr, true)
	return info.handshake(nil)
}

func (c *Conn) tlsInfo(network, addr string, createIfNotFound bool) (info tlsInfo, created bool) {
	// Get info andor create if we're allowed
	c.tlsInfoLock.Lock()
	defer c.tlsInfoLock.Unlock()
	key := network + "!" + addr
	if info = c.tlsInfos[key]; info == nil && createIfNotFound {
		info = newTLSInfo(c, network, addr)
		c.tlsInfos[key] = info
		created = true
	}
	return
}

func (c *Conn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	// TODO: this is inefficient, do better w/ a shared buffer or something
	// TODO: configurable MTU
	var buf [508]byte
	// We continually read until we have read some data. No data and no error
	// means that it wants us to continue, e.g. for a handshake.
	for n == 0 && err == nil {
		if n, addr, err = c.conn.ReadFrom(buf[:]); err != nil {
			n = 0
			break
		}
		if c.client {
			n, err = c.clientReadFrom(buf[:n], addr, p)
		} else {
			n, err = c.serverReadFrom(buf[:n], addr, p)
		}
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
