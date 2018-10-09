package model

const HandshakeTypeHelloRequest HandshakeType = 0

type HandshakeHelloRequest struct{}

func (*HandshakeHelloRequest) Type() HandshakeType { return HandshakeTypeHelloRequest }

func (*HandshakeHelloRequest) Marshal() []byte {
	return []byte{}
}

func (*HandshakeHelloRequest) Unmarshal(b []byte) error {
	if len(b) != 0 {
		return AlertDecodeError
	}
	return nil
}
