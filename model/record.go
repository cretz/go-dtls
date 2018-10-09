package model

type RecordType uint8

const (
	RecordTypeChangeCipherSpec RecordType = 20
	RecordTypeAlert            RecordType = 21
	RecordTypeHandshake        RecordType = 22
	RecordTypeApplicationData  RecordType = 23
)

type DTLSRecord struct {
	Type            RecordType
	ProtocolVersion ProtocolVersion
	Epoch           uint16
	SequenceNumber  uint64
	Fragment        []byte
}

func (d *DTLSRecord) Marshal() ([]byte, error) {
	panic("TODO")
}

func (d *DTLSRecord) Unmarshal(b []byte) (int, error) {
	panic("TODO")
}
