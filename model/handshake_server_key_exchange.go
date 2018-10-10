package model

import (
	"crypto/elliptic"
	"crypto/tls"
	"fmt"
	"io"

	"golang.org/x/crypto/curve25519"
)

const HandshakeTypeServerKeyExchange HandshakeType = 12

type HandshakeServerKeyExchange struct {
	Params       []byte
	SignedParams *DigitallySigned
}

func (*HandshakeServerKeyExchange) Type() HandshakeType { return HandshakeTypeServerKeyExchange }

func (h *HandshakeServerKeyExchange) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeServerKeyExchange) Unmarshal(b []byte) error {
	// TODO: validate sig here?
	panic("TODO")
}

func (h *HandshakeServerKeyExchange) ParseECDHParams() (*ECDHParams, error) {
	// Named curve at https://tools.ietf.org/html/rfc4492#section-5.4
	if len(h.Params) < 4 {
		return nil, fmt.Errorf("Key too small, got %v", h.Params)
	} else if h.Params[0] != 3 {
		return nil, fmt.Errorf("Expected curve 3, got %v", h.Params[0])
	} else if int(h.Params[3])+4 != len(h.Params) {
		return nil, fmt.Errorf("Invalid given length")
	}
	return &ECDHParams{
		CurveID: tls.CurveID(h.Params[1])<<8 | tls.CurveID(h.Params[2]),
		Key:     h.Params[4:],
	}, nil
}

type ECDHParams struct {
	CurveID tls.CurveID
	Key     []byte
}

func (e *ECDHParams) GenerateClientKeyExchange(rand io.Reader) ([]byte, *HandshakeClientKeyExchange, error) {
	// Mostly copied from Go code
	var serialized, preMasterSecret []byte

	if e.CurveID == tls.X25519 {
		var ourPublic, theirPublic, sharedKey, scalar [32]byte

		if _, err := io.ReadFull(rand, scalar[:]); err != nil {
			return nil, nil, err
		}

		copy(theirPublic[:], e.Key)
		curve25519.ScalarBaseMult(&ourPublic, &scalar)
		curve25519.ScalarMult(&sharedKey, &scalar, &theirPublic)
		serialized = ourPublic[:]
		preMasterSecret = sharedKey[:]
	} else {
		curve, ok := curveForCurveID(e.CurveID)
		if !ok {
			return nil, nil, fmt.Errorf("Unrecognized curve ID %v", e.CurveID)
		}
		keyX, keyY := elliptic.Unmarshal(curve, e.Key)
		if keyX == nil {
			return nil, nil, fmt.Errorf("Failed unmarshalling curve key")
		}
		priv, mx, my, err := elliptic.GenerateKey(curve, rand)
		if err != nil {
			return nil, nil, err
		}
		x, _ := curve.ScalarMult(keyX, keyY, priv)
		preMasterSecret = make([]byte, (curve.Params().BitSize+7)>>3)
		xBytes := x.Bytes()
		copy(preMasterSecret[len(preMasterSecret)-len(xBytes):], xBytes)

		serialized = elliptic.Marshal(curve, mx, my)
	}

	ret := &HandshakeClientKeyExchange{ExchangeKeys: make([]byte, 1+len(serialized))}
	ret.ExchangeKeys[0] = byte(len(serialized))
	copy(ret.ExchangeKeys[1:], serialized)

	return preMasterSecret, ret, nil
}

func curveForCurveID(id tls.CurveID) (elliptic.Curve, bool) {
	switch id {
	case tls.CurveP256:
		return elliptic.P256(), true
	case tls.CurveP384:
		return elliptic.P384(), true
	case tls.CurveP521:
		return elliptic.P521(), true
	default:
		return nil, false
	}

}
