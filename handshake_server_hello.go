package dtls

import (
	"context"

	"github.com/cretz/go-dtls/model"
)

type serverHello struct {
	ServerHello        *model.HandshakeServerHello
	Certificate        *model.HandshakeCertificate
	ServerKeyExchange  *model.HandshakeServerKeyExchange
	CertificateRequest *model.HandshakeCertificateRequest
	// Only present if this was a successful resumption
	Finished *model.HandshakeFinished
}

func receiveServerHello(
	ctx context.Context,
	conn AddrConn,
	resumeSessionID []byte,
) (*model.HandshakeHelloRequest, *model.HandshakeHelloVerifyRequest, *serverHello, error) {
	panic("TODO")
}
