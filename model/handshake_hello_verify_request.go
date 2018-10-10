package model

const HandshakeTypeHelloVerifyRequest HandshakeType = 3

type HandshakeHelloVerifyRequest struct {
	ServerVersion uint16
	Cookie        []byte
}

func (*HandshakeHelloVerifyRequest) Type() HandshakeType { return HandshakeTypeHelloVerifyRequest }

func (h *HandshakeHelloVerifyRequest) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeHelloVerifyRequest) Unmarshal(b []byte) error {
	panic("TODO")
}
