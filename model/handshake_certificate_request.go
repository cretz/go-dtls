package model

const HandshakeTypeCertificateRequest HandshakeType = 13

type ClientCertificateType uint8

const (
	ClientCertificateTypeRSASign    ClientCertificateType = 1
	ClientCertificateTypeDSSSign    ClientCertificateType = 2
	ClientCertificateTypeRSAFixedDH ClientCertificateType = 3
	ClientCertificateTypeDSSFixedDH ClientCertificateType = 4
)

type HandshakeCertificateRequest struct {
	CertificateTypes             []ClientCertificateType
	SupportedSignatureAlgorithms []SignatureAndHashAlgorithm
	CertificateAuthorities       []byte
}

func (*HandshakeCertificateRequest) Type() HandshakeType { return HandshakeTypeCertificateRequest }

func (h *HandshakeCertificateRequest) Marshal() []byte {
	panic("TODO")
}

func (h *HandshakeCertificateRequest) Unmarshal(b []byte) error {
	panic("TODO")
}
