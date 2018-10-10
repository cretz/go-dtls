package handshake

import (
	"context"
	"time"
)

func ClientHandshake(ctx context.Context, conn AddrConn) error {
	srvHello, err := clientHandshakeHello(ctx, conn)
	if err == nil {
		err = clientHandshakeFinish(ctx, conn, srvHello)
	}
	return err
}

func clientHandshakeHello(ctx context.Context, conn AddrConn) (*serverHello, error) {
MainLoop:
	// TODO: configurable timeout/interval
	var cookie []byte
	for timeout := time.Second; timeout <= 64*time.Second; timeout *= 2 {
		// Send client hello
		localCtx, localCtxCancelFn := context.WithTimeout(ctx, timeout)
		err := sendClientHello(localCtx, conn, cookie)
		localCtxCancelFn()
		if err == context.DeadlineExceeded {
			// Timeout, just move on
			continue
		} else if err != nil {
			return nil, err
		}

		// Receive hello request, verify hello, or server hello
		localCtx, localCtxCancelFn = context.WithTimeout(ctx, timeout)
		helloReq, verifyReq, srvHello, err := receiveServerHello(localCtx, conn)
		localCtxCancelFn()
		if err == context.DeadlineExceeded {
			// Timeout, just move on
			continue
		} else if err != nil {
			return nil, err
		} else if helloReq != nil {
			// Request start from beginning w/ empty cookie
			cookie = nil
			goto MainLoop
		} else if verifyReq != nil {
			// Request start from beginning with given cookie
			cookie = verifyReq.Cookie
		} else {
			return srvHello, nil
		}
	}
	return nil, context.DeadlineExceeded
}

func clientHandshakeFinish(ctx context.Context, conn AddrConn, srvHello *serverHello) error {
MainLoop:
	for timeout := time.Second; timeout <= 64*time.Second; timeout *= 2 {
		// Send client finish
		localCtx, localCtxCancelFn := context.WithTimeout(ctx, timeout)
		err := sendClientFinish(localCtx, conn, srvHello)
		if err == context.DeadlineExceeded {
			// Timeout, just move on
			continue
		} else if err != nil {
			return err
		}

		// Receive server finish or the server hello again
		localCtx, localCtxCancelFn = context.WithTimeout(ctx, timeout)
		newSrvHello, _, err := receiveServerFinish(localCtx, conn)
		localCtxCancelFn()
		if err == context.DeadlineExceeded {
			// Timeout, just move on
			continue
		} else if err != nil {
			return err
		} else if newSrvHello != nil {
			// Got another server hello, restart from beginning
			srvHello = newSrvHello
			goto MainLoop
		} else {
			return nil
		}
	}
	return context.DeadlineExceeded
}
