package model

const HandshakeTypeServerKeyExchange HandshakeType = 12

type HandshakeServerKeyExchange struct {
	Key []byte
}

func (*HandshakeServerKeyExchange) Type() HandshakeType { return HandshakeTypeServerKeyExchange }

func (h *HandshakeServerKeyExchange) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeServerKeyExchange) Unmarshal(b []byte) error {
	panic("TODO")
}
