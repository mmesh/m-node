package maddr

import (
	"strings"

	"mmesh.dev/m-node/internal/app/node/mnet/p2p/transport"
)

func GetTransport(maddr string) transport.Protocol {
	s := strings.Split(maddr, "/")

	if len(s) < 5 {
		return transport.Invalid
	}

	proto := transport.Protocol(strings.ToUpper(s[3]))

	switch proto {
	case transport.ProtocolTCP:
		return transport.ProtocolTCP
	case transport.ProtocolUDP:
		if transport.Protocol(strings.ToUpper(s[5])) == transport.ProtocolQUIC {
			return transport.ProtocolQUIC
		}
	}

	return transport.Invalid
}
