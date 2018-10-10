package model

const HandshakeTypeClientHello HandshakeType = 1

const CompressionMethodNone uint8 = 0

type HandshakeClientHello struct {
	ClientVersion      uint16
	Random             []byte
	SessionID          []byte
	Cookie             []byte
	CipherSuites       []uint16
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
