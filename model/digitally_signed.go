package model

type HashAlgorithm uint8

const (
	HashAlgorithmNone   HashAlgorithm = 0
	HashAlgorithmMD5    HashAlgorithm = 1
	HashAlgorithmSHA1   HashAlgorithm = 2
	HashAlgorithmSHA224 HashAlgorithm = 3
	HashAlgorithmSHA256 HashAlgorithm = 4
	HashAlgorithmSHA384 HashAlgorithm = 5
	HashAlgorithmSHA512 HashAlgorithm = 6
)

type SignatureAlgorithm uint8

const (
	SignatureAlgorithmAnonymous SignatureAlgorithm = 0
	SignatureAlgorithmRSA       SignatureAlgorithm = 1
	SignatureAlgorithmDSA       SignatureAlgorithm = 2
	SignatureAlgorithmECDSA     SignatureAlgorithm = 3
)

type SignatureAndHashAlgorithm uint16

func (s SignatureAndHashAlgorithm) HashAlgorithm() HashAlgorithm {
	return HashAlgorithm(s >> 8)
}
func (s SignatureAndHashAlgorithm) SignatureAlgorithm() SignatureAlgorithm {
	return SignatureAlgorithm(s)
}

type DigitallySigned struct {
	Algorithm SignatureAndHashAlgorithm
	Signature []byte
}
