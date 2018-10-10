package dtls

import (
	"context"
	"crypto/x509"
	"time"

	"github.com/cretz/go-dtls/model"
)

type clientHandshakeState struct {
	*addrConn

	resumeSessionID []byte
	cookie          []byte

	clientHello *model.HandshakeClientHello

	serverHello              *model.HandshakeServerHello
	serverCertificate        *model.HandshakeCertificate
	serverKeyExchange        *model.HandshakeServerKeyExchange
	serverCertificateRequest *model.HandshakeCertificateRequest
	serverFinished           bool
	serverECDHParams         *model.ECDHParams
	serverCertificates       []*x509.Certificate

	masterSecret      []byte
	clientKeyExchange *model.HandshakeClientKeyExchange
}

func clientHandshake(ctx context.Context, conn *addrConn, resumeSessionID []byte) error {
	hs := &clientHandshakeState{addrConn: conn, resumeSessionID: resumeSessionID}
	return hs.handshake(ctx)
}

func (hs *clientHandshakeState) handshake(ctx context.Context) error {
	// TODO: eager validation of the conf here (e.g. cipher suite)? or where addrConn was created?
	err := hs.hello(ctx)
	if err == nil {
		err = hs.finish(ctx)
	}
	return err
}

func (hs *clientHandshakeState) hello(ctx context.Context) error {
MainLoop:
	// TODO: configurable timeout/interval
	for timeout := time.Second; timeout <= 64*time.Second; timeout *= 2 {
		// Send client hello
		if err := hs.withTimeout(ctx, timeout, hs.handshakeClientHello); err == context.DeadlineExceeded {
			// Timeout, just move on
			continue
		} else if err != nil {
			return err
		}

		// Receive hello request, verify hello, or server hello
		restart := false
		err := hs.withTimeout(ctx, timeout, func(ctx context.Context) (err error) {
			restart, err = hs.handshakeServerHello(ctx)
			return
		})
		if err == context.DeadlineExceeded {
			// Timeout, just move on
			continue
		} else if err != nil {
			return err
		} else if restart {
			// TODO: prevent infinite restarts
			goto MainLoop
		} else {
			return nil
		}
	}
	return context.DeadlineExceeded
}

func (hs *clientHandshakeState) finish(ctx context.Context) error {
MainLoop:
	for timeout := time.Second; timeout <= 64*time.Second; timeout *= 2 {
		// Send client finish
		if err := hs.withTimeout(ctx, timeout, hs.handshakeClientFinish); err == context.DeadlineExceeded {
			// Timeout, just move on
			continue
		} else if err != nil {
			return err
		}

		// If server finish already happened, this is enough...
		if hs.serverFinished {
			return nil
		}

		// Receive server finish or the server hello again
		restart := false
		err := hs.withTimeout(ctx, timeout, func(ctx context.Context) (err error) {
			restart, err = hs.handshakeServerFinish(ctx)
			return
		})
		if err == context.DeadlineExceeded {
			// Timeout, just move on
			continue
		} else if err != nil {
			return err
		} else if restart {
			// TODO: prevent infinite restarts
			goto MainLoop
		} else {
			return nil
		}
	}
	return context.DeadlineExceeded
}
