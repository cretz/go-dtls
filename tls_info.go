package dtls

import (
	"context"
	"time"

	"github.com/cretz/go-dtls/model"
)

type tlsInfo interface {
	// Record can be nil to start client handshake
	handshake(rec *model.DTLSRecord) error
}

func newTLSInfo(conn *Conn, network, addr string) tlsInfo {
	if conn.client {
		return &clientTLSInfo{conn, network, addr}
	}
	panic("TODO: server")
}

type clientTLSInfo struct {
	conn    *Conn
	network string
	addr    string
}

func (c *clientTLSInfo) handshake(rec *model.DTLSRecord) error {
	// Don't run if established? but what about hello request?
	panic("TODO")
}

func (c *clientTLSInfo) established() bool {
	panic("TODO")
}

func (c *clientTLSInfo) doHandshake(ctx context.Context) error {
	/*
		TODO:

		* Send client hello
		* Wait for either: 1) hello request, 2) hello verify request, or 3) server hello + stuff + server hello done
	*/
	var serverInfo interface{} = nil
	err := c.doFlight(ctx, c.sendClientHello, func(ctx context.Context) (restart bool, err error) {
		var serverInfo interface{} = nil

	})
}

func (c *clientTLSInfo) sendClientHello(ctx context.Context) error {
	panic("TODO")
}

type serverHello struct {
	ServerHello        *model.HandshakeServerHello
	Certificate        *model.HandshakeCertificate
	ServerKeyExchange  *model.HandshakeServerKeyExchange
	CertificateRequest *model.HandshakeCertificateRequest
}

// All nil just means restart
func (c *clientTLSInfo) receiveServerHello(ctx context.Context) (*model.HandshakeHelloVerifyRequest, *serverHello, error) {
	var hello *serverHello
	// TODO: lots of validation including message seq, order, etc
	for {
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		case h := <-c.nextHandshakeMessage():
			if h.Err != nil {
				return nil, nil, h.Err
			}
			body, err := h.Msg.Body()
			if err != nil {
				return nil, nil, err
			}
			switch body := body.(type) {
			case nil:
				return nil, nil, h.Err
			case *model.HandshakeHelloRequest:
				// Retransmission request means start over
				return nil, nil, nil
			case *model.HandshakeHelloVerifyRequest:
				return body, nil, nil
			case *model.HandshakeServerHello:
				hello := &serverHello{}
			case *model.HandshakeCertificate:
				if hello == nil {
					return nil, nil, model.AlertUnexpectedMessage
				}
				hello.Certificate = body
			case *model.HandshakeServerKeyExchange:
				if hello == nil {
					return nil, nil, model.AlertUnexpectedMessage
				}
				panic("TODO")
			case *model.HandshakeCertificateRequest:

			}
		}
	}
}

type handshakeOrErr struct {
	Err error
	Msg *model.Handshake
}

func (c *clientTLSInfo) nextHandshakeMessage() <-chan *handshakeOrErr {
	panic("TODO")
}

func (c *clientTLSInfo) doFlight(
	ctx context.Context,
	send func(context.Context) error,
	receive func(context.Context) (restart bool, err error),
) error {
	ctx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()
RestartLoop:
	for {
		// TODO: time configurable
		for i := time.Second; i <= 64; i *= 2 {
			// Send first
			attemptCtx, attemptCancelFn := context.WithTimeout(ctx, i)
			err := send(attemptCtx)
			attemptCancelFn()
			if err != nil && err != context.DeadlineExceeded {
				return err
			}
			// Now try to receive
			attemptCtx, attemptCancelFn = context.WithTimeout(ctx, i)
			restart, err := receive(attemptCtx)
			attemptCancelFn()
			if err == nil {
				if restart {
					continue RestartLoop
				}
				return nil
			} else if err != context.DeadlineExceeded {
				return err
			}
		}
	}
}
