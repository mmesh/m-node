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

	proto := transport.Protocol(strings.ToLower(s[3]))

	switch proto {
	case transport.ProtocolTCP:
		return transport.ProtocolTCP
	case transport.ProtocolUDP:
		switch transport.Protocol(strings.ToLower(s[5])) {
		case transport.ProtocolQUIC:
			return transport.ProtocolQUIC
		case transport.ProtocolQUICv1:
			return transport.ProtocolQUICv1
		}
	}

	return transport.Invalid
}
