package model

const HandshakeTypeCertificateVerify HandshakeType = 15

type HandshakeCertificateVerify struct {
	HandshakeMessages    []byte
	HandshakeMessagesSig *DigitallySigned
}

func (*HandshakeCertificateVerify) Type() HandshakeType { return HandshakeTypeCertificateVerify }

func (h *HandshakeCertificateVerify) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeCertificateVerify) Unmarshal(b []byte) error {
	panic("TODO")
}
