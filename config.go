package dtls

import (
	"crypto/rand"
	"crypto/tls"
	"io"

	"github.com/cretz/go-dtls/model"
)

const (
	VersionDTLS10 = 0xFEFF
	VersionDTLS12 = 0xFEFD
)

type Config struct {
	// Values that are used:
	//
	// * Rand
	// * CipherSuites
	//
	// All other values are ignored
	*tls.Config
}

func (c *Config) rand() io.Reader {
	r := c.Rand
	if r == nil {
		return rand.Reader
	}
	return r
}

func (c *Config) cipherSuites() []uint16 {
	s := c.CipherSuites
	if s == nil {
		s = model.CipherSuitesDefault
	}
	return s
}
