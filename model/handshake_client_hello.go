package model

const HandshakeTypeClientHello HandshakeType = 1

type HandshakeClientHello struct {
	ClientVersion      ProtocolVersion
	Random             HandshakeRandom
	SessionID          []byte
	Cookie             []byte
	CipherSuites       []CipherSuite
	CompressionMethods []uint8
	Extensions         []Extension
}

func (*HandshakeClientHello) Type() HandshakeType { return HandshakeTypeClientHello }

func (h *HandshakeClientHello) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeClientHello) Unmarshal(b []byte) error {
	panic("TODO")
}
