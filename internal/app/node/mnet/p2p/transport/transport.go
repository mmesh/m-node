package transport

type Protocol string

const (
	ProtocolUDP  Protocol = "UDP"
	ProtocolTCP  Protocol = "TCP"
	ProtocolQUIC Protocol = "QUIC"

	Invalid Protocol = "-"
)

func (p Protocol) String() string {
	return string(p)
}
