package transport

type Protocol string

const (
	ProtocolUDP    Protocol = "udp"
	ProtocolTCP    Protocol = "tcp"
	ProtocolQUIC   Protocol = "quic"
	ProtocolQUICv1 Protocol = "quic-v1"

	Invalid Protocol = "-"
)

func (p Protocol) String() string {
	return string(p)
}
