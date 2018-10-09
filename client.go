package dtls

import (
	"net"

	"github.com/cretz/go-dtls/model"
)

// If result is 0, nil then callers should read more immediately
func (c *Conn) clientReadFrom(b []byte, addr net.Addr, p []byte) (int, error) {
	// Always require an entry for the server
	info, _ := c.tlsInfo(addr.Network(), addr.String(), false)
	if info == nil {
		// TODO: can this be used as a timing attack by others to see what servers
		// I might be connected to?
		return 0, &ErrNoServerConnEstablished{addr, b}
	}
	// Packet can contain multiple messages
	for len(b) > 0 {
		// Parse the message
		rec := &model.DTLSRecord{}
		n, err := rec.Unmarshal(b)
		if err != nil {
			return 0, err
		}
		b = b[:n]
		// TODO: validate message, e.g. version, size, etc
		// Handle message types
		switch rec.Type {
		case model.RecordTypeChangeCipherSpec:
			panic("TODO")
		case model.RecordTypeAlert:
			panic("TODO")
		case model.RecordTypeHandshake:
			if err = info.handshake(rec); err != nil {
				return 0, err
			}
		case model.RecordTypeApplicationData:
			panic("TODO")
		default:
			return 0, model.AlertUnexpectedMessage
		}
	}
	// This read did nothing
	return 0, nil
}
