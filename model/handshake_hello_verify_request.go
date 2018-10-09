package model

const HandshakeTypeHelloVerifyRequest HandshakeType = 3

type HandshakeHelloVerifyRequest struct {
	ServerVersion ProtocolVersion
	Cookie        []byte
}

func (*HandshakeHelloVerifyRequest) Type() HandshakeType { return HandshakeTypeHelloVerifyRequest }

func (h *HandshakeHelloVerifyRequest) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeHelloVerifyRequest) Unmarshal(b []byte) error {
	panic("TODO")
}
