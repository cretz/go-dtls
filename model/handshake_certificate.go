package model

import (
	"crypto/x509"
)

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

// Error is not a bad certificate alert, caller is expected to handle that
func (h *HandshakeCertificate) ParseX509Certificates() (certs []*x509.Certificate, err error) {
	certs = make([]*x509.Certificate, len(h.CertificateList))
	for i, certData := range h.CertificateList {
		if certs[i], err = x509.ParseCertificate(certData); err != nil {
			return nil, err
		}
	}
	return
}
