package model

const HandshakeTypeCertificateRequest HandshakeType = 13

type HandshakeCertificateRequest struct {
}

func (*HandshakeCertificateRequest) Type() HandshakeType { return HandshakeTypeCertificateRequest }

func (h *HandshakeCertificateRequest) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeCertificateRequest) Unmarshal(b []byte) error {
	panic("TODO")
}
