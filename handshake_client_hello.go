package dtls

import (
	"context"
	"fmt"
	"io"

	"github.com/cretz/go-dtls/model"
)

func (hs *clientHandshakeState) handshakeClientHello(ctx context.Context) error {
	hs.clientHello = &model.HandshakeClientHello{
		ClientVersion:      VersionDTLS12,
		Random:             make([]byte, 32),
		SessionID:          hs.resumeSessionID,
		Cookie:             hs.cookie,
		CipherSuites:       hs.config.cipherSuites(),
		CompressionMethods: []uint8{model.CompressionMethodNone},
	}
	// Go TLS doesn't put current time in random, so we aren't gonna either
	if _, err := io.ReadFull(hs.config.rand(), hs.clientHello.Random); err != nil {
		return fmt.Errorf("Failed reading rand: %v", err)
	}
	// TODO: extensions
	return hs.sendHandshake(ctx, hs.clientHello)
}
