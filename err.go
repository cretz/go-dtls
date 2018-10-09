package dtls

import (
	"fmt"
	"net"
)

type ErrNoServerConnEstablished struct {
	Addr   net.Addr
	Packet []byte
}

func (e *ErrNoServerConnEstablished) Error() string {
	return fmt.Sprintf("No established connection with %v", e.Addr)
}
