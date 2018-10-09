package model

type ProtocolVersion uint16

func (p ProtocolVersion) Major() uint8 { return uint8(p >> 8) }
func (p ProtocolVersion) Minor() uint8 { return uint8(p) }

type CipherSuite uint16
