package dtls

import (
	"bytes"
	"context"

	"github.com/cretz/go-dtls/model"
)

func (hs *clientHandshakeState) handshakeServerHello(ctx context.Context) (restart bool, err error) {
	// Clear out pieces that may have happened before
	hs.serverHello = nil
	hs.serverCertificate = nil
	hs.serverKeyExchange = nil
	hs.serverCertificateRequest = nil
	hs.serverFinished = false
	hs.serverECDHParams = nil
	hs.serverCertificates = nil

	// Get server hello or a restart request
	if restart, err = hs.receiveServerHello(ctx); restart || err != nil {
		return
	}

	// TODO: extensions

	// Exit early if it's a resumption
	if len(hs.resumeSessionID) > 0 && bytes.Equal(hs.resumeSessionID, hs.serverHello.SessionID) {
		return hs.handshakeServerFinish(ctx)
	}

	// We always expect a certificate chain when not resuming
	// TODO: DH_anon support someday?
	// TODO: validate the leaf wasn't changed during renegotiation per Go TLS src
	var ok bool
	if _, body, err := hs.receiveHandshake(ctx); err != nil {
		return false, err
	} else if hs.serverCertificate, ok = body.(*model.HandshakeCertificate); !ok {
		return false, hs.alertf(model.AlertUnexpectedMessage, "Expected cert, got %v", body)
	}
	certs, err := hs.serverCertificate.ParseX509Certificates()
	if err != nil {
		return false, hs.alertf(model.AlertBadCertificate, "Failed parsing certs: %v", err)
	}
	// TODO: validate certs
	hs.serverCertificates = certs

	// TODO: OCSP support w/ certificate status

	// Some of the next items are optional
	_, body, err := hs.receiveHandshake(ctx)
	if err != nil {
		return false, err
	}
	// Handle key exchange info
	if hs.serverKeyExchange, ok = body.(*model.HandshakeServerKeyExchange); ok {
		if hs.serverECDHParams, err = hs.serverKeyExchange.ParseECDHParams(); err != nil {
			return false, hs.alertf(model.AlertUnexpectedMessage, "Failed parsing ECDH params: %v", err)
		}
		if _, body, err = hs.receiveHandshake(ctx); err != nil {
			return false, err
		}
	}
	// Handle certificate request
	if hs.serverCertificateRequest, ok = body.(*model.HandshakeCertificateRequest); ok {
		if _, body, err = hs.receiveHandshake(ctx); err != nil {
			return false, err
		}
	}
	// Check for done
	if _, ok := body.(*model.HandshakeServerHelloDone); !ok {
		return false, hs.alertf(model.AlertUnexpectedMessage, "Expected server done, got %v", body)
	}
	return false, nil
}

func (hs *clientHandshakeState) receiveServerHello(ctx context.Context) (restart bool, err error) {
	// Hello req, verify req, server hello, or fail
	_, body, err := hs.receiveHandshake(ctx)
	if err != nil {
		return false, err
	}
	switch body := body.(type) {
	case *model.HandshakeHelloRequest:
		// Remove cookie and restart
		// TODO: disable renegotiation support? disable mid-handshake?
		hs.cookie = nil
		return true, nil
	case *model.HandshakeHelloVerifyRequest:
		if body.ServerVersion != VersionDTLS10 {
			return false, hs.alertf(model.AlertIllegalParameter,
				"Expected version %v on verify, got %v", VersionDTLS10, body.ServerVersion)
		}
		// Update cookie and restart
		hs.cookie = body.Cookie
		return true, nil
	case *model.HandshakeServerHello:
		// Validation occurs below
		hs.serverHello = body
	default:
		return false, hs.alertf(model.AlertUnexpectedMessage, "Expecting server hello, got %v", body)
	}
	// Validate the server hello
	// We only do 1.2 for now
	if hs.serverHello.ServerVersion != VersionDTLS12 {
		return false, hs.alertf(model.AlertProtocolVersion,
			"Expecting version %v, server hello gave %v", VersionDTLS12, hs.serverHello.ServerVersion)
	}
	// The cipher suite must be something we passed in
	// TODO: validate this wasn't changed on renegotiation
	foundCipherSuite := false
	for _, cipherSuite := range hs.clientHello.CipherSuites {
		if cipherSuite == hs.serverHello.CipherSuite {
			foundCipherSuite = true
			break
		}
	}
	if !foundCipherSuite {
		return false, hs.alertf(model.AlertHandshakeFailure,
			"No server suite %v in client set %v", hs.serverHello.CipherSuite, hs.clientHello.CipherSuites)
	}
	// No compression method
	if hs.serverHello.CompressionMethod != model.CompressionMethodNone {
		return false, hs.alertf(model.AlertHandshakeFailure,
			"Expected no compression method, got %v", hs.serverHello.CompressionMethod)
	}
	return false, nil
}
