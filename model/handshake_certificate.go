package model

const HandshakeTypeCertificate HandshakeType = 11

type HandshakeCertificate struct {
	CertificateList [][]byte
}

func (*HandshakeCertificate) Type() HandshakeType { return HandshakeTypeCertificate }

func (h *HandshakeCertificate) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeCertificate) Unmarshal(b []byte) error {
	panic("TODO")
}
