package model

import "time"

type HandshakeType uint8

const (
// HandshakeTypeHelloRequest       HandshakeType = 0
// HandshakeTypeClientHello        HandshakeType = 1
// HandshakeTypeServerHello        HandshakeType = 2
// HandshakeTypeHelloVerifyRequest HandshakeType = 3
// HandshakeTypeCertificate        HandshakeType = 11
// HandshakeTypeServerKeyExchange  HandshakeType = 12
// HandshakeTypeCertificateRequest HandshakeType = 13
// HandshakeTypeServerHelloDone    HandshakeType = 14
// HandshakeTypeCertificateVerify  HandshakeType = 15
// HandshakeTypeClientKeyExchange  HandshakeType = 16
// HandshakeTypeFinished           HandshakeType = 20
)

var HandshakeBodyCreators = map[HandshakeType]func() HandshakeBody{
	HandshakeTypeHelloRequest:       func() HandshakeBody { return &HandshakeHelloRequest{} },
	HandshakeTypeClientHello:        func() HandshakeBody { return &HandshakeClientHello{} },
	HandshakeTypeServerHello:        func() HandshakeBody { return &HandshakeServerHello{} },
	HandshakeTypeHelloVerifyRequest: func() HandshakeBody { return &HandshakeHelloVerifyRequest{} },
	HandshakeTypeCertificate:        func() HandshakeBody { return &HandshakeCertificate{} },
	HandshakeTypeServerKeyExchange:  func() HandshakeBody { return &HandshakeServerKeyExchange{} },
	HandshakeTypeCertificateRequest: func() HandshakeBody { return &HandshakeCertificateRequest{} },
}

type HandshakeFragment struct {
	Type       HandshakeType
	MessageSeq uint16
	// Ignored except on the parameter to Merge
	FragmentOffset uint32
	Bytes          []byte
	TotalLength    uint16
}

// Zeros out the current fragment's offset. Doesn't mutate other.
func (h *HandshakeFragment) Merge(other *HandshakeFragment) error {
	panic("TODO")
}

// Returns nil when not complete
func (h *HandshakeFragment) Complete() *Handshake {
	panic("TODO")
}

type Handshake struct {
	Type       HandshakeType
	MessageSeq uint16
	RawBody    []byte
}

// Seq is not set here
func HandshakeFromBody(body HandshakeBody) (*Handshake, error) {
	panic("TODO")
}

func (h *Handshake) Body() (HandshakeBody, error) {
	panic("TODO")
}

type HandshakeRandom struct {
	Time        time.Time
	RandomBytes []byte
}

type HandshakeBody interface {
	Type() HandshakeType
	Marshal() []byte
	Unmarshal([]byte) error
}
