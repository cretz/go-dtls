package dtls

import (
	"io"
	"time"
)

type Config struct {
	Rand       io.Reader
	Time       func() time.Time
	MinVersion uint16
	MaxVersion uint16
}
