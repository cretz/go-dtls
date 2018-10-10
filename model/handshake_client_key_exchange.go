package model

const HandshakeTypeClientKeyExchange HandshakeType = 16

type HandshakeClientKeyExchange struct {
	ExchangeKeys []byte
}

func (*HandshakeClientKeyExchange) Type() HandshakeType { return HandshakeTypeClientKeyExchange }

func (h *HandshakeClientKeyExchange) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeClientKeyExchange) Unmarshal(b []byte) error {
	panic("TODO")
}
