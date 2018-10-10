package model

const HandshakeTypeServerHello HandshakeType = 2

type HandshakeServerHello struct {
	ServerVersion     uint16
	Random            []byte
	SessionID         []byte
	CipherSuite       uint16
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
