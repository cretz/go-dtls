package dtls

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"github.com/cretz/go-dtls/model"
)

func (hs *clientHandshakeState) handshakeClientFinish(ctx context.Context) (err error) {
	// Clear out anything set before
	hs.masterSecret = nil
	hs.clientKeyExchange = nil

	var preMasterSecret []byte
	// If the server has already finished, skip some steps
	if !hs.serverFinished {
		// Requested cert? Send it...
		if hs.serverCertificateRequest != nil {
			if cert, err := hs.getCertificate(ctx); err != nil {
				return err
			} else if err = hs.sendHandshake(ctx, &model.HandshakeCertificate{cert.Certificate}); err != nil {
				return err
			}
		}
		// Send client key exchange
		if hs.serverECDHParams == nil {
			return hs.alertf(model.AlertUnexpectedMessage, "Missing server key exchange")
		}
		preMasterSecret, hs.clientKeyExchange, err = hs.serverECDHParams.GenerateClientKeyExchange(hs.config.rand())
		if err != nil {
			return
		} else if err = hs.sendHandshake(ctx, hs.clientKeyExchange); err != nil {
			return
		}
	}
	panic("TODO")
}

func (hs *clientHandshakeState) getCertificate(ctx context.Context) (*tls.Certificate, error) {
	// Use callback if present
	if getCertFn := hs.config.GetClientCertificate; getCertFn != nil {
		reqInfo := &tls.CertificateRequestInfo{
			AcceptableCAs:    hs.serverCertificateRequest.CertificateAuthorities,
			SignatureSchemes: make([]tls.SignatureScheme, len(hs.serverCertificateRequest.SupportedSignatureAlgorithms)),
		}
		for i, sig := range hs.serverCertificateRequest.SupportedSignatureAlgorithms {
			reqInfo.SignatureSchemes[i] = tls.SignatureScheme(sig)
		}
		return getCertFn(reqInfo)
	}

	// Rest is mostly copied from Go source...
	var rsaAvail, ecdsaAvail bool
	for _, certType := range hs.serverCertificateRequest.CertificateTypes {
		switch certType {
		case model.ClientCertificateTypeRSASign:
			rsaAvail = true
		case model.ClientCertificateTypeECDSASign:
			ecdsaAvail = true
		}
	}
	if !rsaAvail && !ecdsaAvail {
		return &tls.Certificate{}, nil
	}
findCert:
	for i, chain := range hs.config.Certificates {
		for j, cert := range chain.Certificate {
			// Leaf as an optimization on first, otherwise parse
			x509Cert := chain.Leaf
			if j != 0 || x509Cert == nil {
				var err error
				if x509Cert, err = x509.ParseCertificate(cert); err != nil {
					return nil, fmt.Errorf("Unable to parse cert %v: %v", i, err)
				}
			}
			switch {
			case rsaAvail && x509Cert.PublicKeyAlgorithm == x509.RSA:
			case ecdsaAvail && x509Cert.PublicKeyAlgorithm == x509.ECDSA:
			default:
				continue findCert
			}
			// No CA, give first
			if len(hs.serverCertificateRequest.CertificateAuthorities) == 0 {
				return &chain, nil
			}
			for _, ca := range hs.serverCertificateRequest.CertificateAuthorities {
				if bytes.Equal(x509Cert.RawIssuer, ca) {
					return &chain, nil
				}
			}
		}
	}
	return &tls.Certificate{}, nil
}
