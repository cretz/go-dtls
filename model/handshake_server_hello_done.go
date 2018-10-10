package model

const HandshakeTypeServerHelloDone HandshakeType = 14

type HandshakeServerHelloDone struct{}

func (*HandshakeServerHelloDone) Type() HandshakeType { return HandshakeTypeServerHelloDone }

func (*HandshakeServerHelloDone) Marshal() []byte {
	return []byte{}
}

func (*HandshakeServerHelloDone) Unmarshal(b []byte) error {
	if len(b) != 0 {
		return AlertDecodeError
	}
	return nil
}
