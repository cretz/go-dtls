package handshake

import (
	"context"

	"github.com/cretz/go-dtls/model"
)

type serverHello struct {
	ServerHello        *model.HandshakeServerHello
	Certificate        *model.HandshakeCertificate
	ServerKeyExchange  *model.HandshakeServerKeyExchange
	CertificateRequest *model.HandshakeCertificateRequest
}

func receiveServerHello(
	ctx context.Context,
	conn AddrConn,
) (*model.HandshakeHelloRequest, *model.HandshakeHelloVerifyRequest, *serverHello, error) {
	panic("TODO")
}
