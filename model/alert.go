package model

// Taken from go/src/crypto/tls/alert.go and modified

import "strconv"

type Alert uint8

const (
	// Alert level
	AlertLevelWarning = 1
	AlertLevelError   = 2
)

const (
	AlertCloseNotify            Alert = 0
	AlertUnexpectedMessage      Alert = 10
	AlertBadRecordMAC           Alert = 20
	AlertDecryptionFailed       Alert = 21
	AlertRecordOverflow         Alert = 22
	AlertDecompressionFailure   Alert = 30
	AlertHandshakeFailure       Alert = 40
	AlertBadCertificate         Alert = 42
	AlertUnsupportedCertificate Alert = 43
	AlertCertificateRevoked     Alert = 44
	AlertCertificateExpired     Alert = 45
	AlertCertificateUnknown     Alert = 46
	AlertIllegalParameter       Alert = 47
	AlertUnknownCA              Alert = 48
	AlertAccessDenied           Alert = 49
	AlertDecodeError            Alert = 50
	AlertDecryptError           Alert = 51
	AlertProtocolVersion        Alert = 70
	AlertInsufficientSecurity   Alert = 71
	AlertInternalError          Alert = 80
	AlertInappropriateFallback  Alert = 86
	AlertUserCanceled           Alert = 90
	AlertNoRenegotiation        Alert = 100
	AlertNoApplicationProtocol  Alert = 120
)

var alertText = map[Alert]string{
	AlertCloseNotify:            "close notify",
	AlertUnexpectedMessage:      "unexpected message",
	AlertBadRecordMAC:           "bad record MAC",
	AlertDecryptionFailed:       "decryption failed",
	AlertRecordOverflow:         "record overflow",
	AlertDecompressionFailure:   "decompression failure",
	AlertHandshakeFailure:       "handshake failure",
	AlertBadCertificate:         "bad certificate",
	AlertUnsupportedCertificate: "unsupported certificate",
	AlertCertificateRevoked:     "revoked certificate",
	AlertCertificateExpired:     "expired certificate",
	AlertCertificateUnknown:     "unknown certificate",
	AlertIllegalParameter:       "illegal parameter",
	AlertUnknownCA:              "unknown certificate authority",
	AlertAccessDenied:           "access denied",
	AlertDecodeError:            "error decoding message",
	AlertDecryptError:           "error decrypting message",
	AlertProtocolVersion:        "protocol version not supported",
	AlertInsufficientSecurity:   "insufficient security level",
	AlertInternalError:          "internal error",
	AlertInappropriateFallback:  "inappropriate fallback",
	AlertUserCanceled:           "user canceled",
	AlertNoRenegotiation:        "no renegotiation",
	AlertNoApplicationProtocol:  "no application protocol",
}

func (e Alert) Text() string {
	return alertText[e]
}

func (e Alert) String() string {
	s, ok := alertText[e]
	if ok {
		return "tls: " + s
	}
	return "tls: Alert(" + strconv.Itoa(int(e)) + ")"
}

func (e Alert) Error() string {
	return e.String()
}
