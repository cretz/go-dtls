package model

const HandshakeTypeServerHello HandshakeType = 2

type HandshakeServerHello struct {
	ServerVersion     ProtocolVersion
	Random            HandshakeRandom
	SessionID         []byte
	CipherSuite       CipherSuite
	CompressionMethod uint8
	Extensions        []Extension
}

func (*HandshakeServerHello) Type() HandshakeType { return HandshakeTypeServerHello }

func (h *HandshakeServerHello) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeServerHello) Unmarshal(b []byte) error {
	panic("TODO")
}
