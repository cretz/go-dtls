package dtls

import (
	"context"
	"time"

	"github.com/cretz/go-dtls/model"
)

type addrConn struct {
	config *Config
}

func (a *addrConn) sendHandshake(ctx context.Context, body model.HandshakeBody) error {
	panic("TODO")
}

func (a *addrConn) sendRecord(ctx context.Context, typ model.RecordType, data []byte) error {
	panic("TODO")
}

func (a *addrConn) receiveHandshake(ctx context.Context) (changeCipherSpec bool, body model.HandshakeBody, err error) {
	// TODO: all but the beginning want to ignore hello requests here btw
	panic("TODO")
}
func (a *addrConn) withTimeout(ctx context.Context, timeout time.Duration, fn func(context.Context) error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	ctx, cancelFn := context.WithTimeout(ctx, timeout)
	defer cancelFn()
	return fn(ctx)
}

func (a *addrConn) alertf(alert model.Alert, format string, v ...interface{}) error {
	panic("TODO")
}
