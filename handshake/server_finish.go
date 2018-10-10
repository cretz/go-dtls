package handshake

import (
	"context"

	"github.com/cretz/go-dtls/model"
)

func receiveServerFinish(ctx context.Context, conn AddrConn) (*serverHello, *model.HandshakeFinished, error) {
	panic("TODO")
}
