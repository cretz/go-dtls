package model

const HandshakeTypeFinished HandshakeType = 20

type HandshakeFinished struct {
	VerifyData []byte
}

func (*HandshakeFinished) Type() HandshakeType { return HandshakeTypeFinished }

func (h *HandshakeFinished) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeFinished) Unmarshal(b []byte) error {
	panic("TODO")
}
