package peer

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-node/internal/app/node/mnet/maddr"
	"mmesh.dev/m-node/internal/app/node/mnet/p2p/transport"
)

// const maxIDChars int = 8

type Hop struct {
	ID      string
	MAddrs  []string
	IsRelay bool
}

type Connection struct {
	PeerInfo         *peer.AddrInfo
	Proto            transport.Protocol
	DirectConnection bool
}

type peerNetHopMap struct {
	direct *peerTransportMap
	relay  *peerTransportMap
}

type peerTransportMap struct {
	quic map[peer.ID]*peer.AddrInfo
	tcp  map[peer.ID]*peer.AddrInfo
}

func newPeerTransportMap() *peerTransportMap {
	return &peerTransportMap{
		quic: make(map[peer.ID]*peer.AddrInfo, 0),
		tcp:  make(map[peer.ID]*peer.AddrInfo, 0),
	}
}

func (ptm *peerTransportMap) update(proto transport.Protocol, pi *peer.AddrInfo) {
	switch proto {
	case transport.ProtocolQUIC:
		if _, ok := ptm.quic[pi.ID]; !ok {
			ptm.quic[pi.ID] = pi
		} else {
			ptm.quic[pi.ID].Addrs = append(ptm.quic[pi.ID].Addrs, pi.Addrs...)
		}
	case transport.ProtocolTCP:
		if _, ok := ptm.tcp[pi.ID]; !ok {
			ptm.tcp[pi.ID] = pi
		} else {
			ptm.tcp[pi.ID].Addrs = append(ptm.tcp[pi.ID].Addrs, pi.Addrs...)
		}
	}
}

func (ptm *peerTransportMap) connect(p2pHost host.Host) (*Connection, error) {
	// QUIC transport
	for _, peerInfo := range ptm.quic {
		// peerID := peerInfo.ID.Pretty()[:maxIDChars]
		// proto := transport.ProtocolQUIC.String()

		// xlog.Debugf("Trying connection to peer %s via %s", peerID, proto)
		if err := connect(p2pHost, peerInfo); err != nil {
			// xlog.Warnf("Unable to connect peer %s via %s: %v", peerID, proto, err)
			continue
		} else {
			// xlog.Debugf("CONNECTED to peer %s via %s", peerID, proto)
			return &Connection{
				PeerInfo: peerInfo,
				Proto:    transport.ProtocolQUIC,
			}, nil
		}
	}

	// TCP transport
	for _, peerInfo := range ptm.tcp {
		// peerID := peerInfo.ID.Pretty()[:maxIDChars]
		// proto := transport.ProtocolTCP.String()

		// xlog.Debugf("Trying connection to peer %s via %s", peerID, proto)
		if err := connect(p2pHost, peerInfo); err != nil {
			// xlog.Warnf("Unable to connect peer %s via %s: %v", peerID, proto, err)
			continue
		}
		// xlog.Debugf("CONNECTED to peer %s via %s", peerID, proto)
		return &Connection{
			PeerInfo: peerInfo,
			Proto:    transport.ProtocolTCP,
		}, nil
	}

	return nil, fmt.Errorf("unable to connect to peer")
}

func getPeerInfoMapFromNetHop(hop *Hop) *peerNetHopMap {
	pnhm := &peerNetHopMap{
		direct: newPeerTransportMap(),
		relay:  newPeerTransportMap(),
	}

	for _, ma := range hop.MAddrs {
		proto := maddr.GetTransport(ma)

		if proto == transport.Invalid {
			continue
		}

		pi, err := getPeerInfo(ma)
		if err != nil {
			// xlog.Errorf("Unable to parse multiaddr %s: %v", ma, err)
			continue
		}

		if maddr.IsRelay(ma) {
			pnhm.relay.update(proto, pi)
		} else {
			pnhm.direct.update(proto, pi)
		}
	}

	return pnhm
}

func ConnectHop(p2pHost host.Host, hop *Hop) (*Connection, error) {
	pnhm := getPeerInfoMapFromNetHop(hop)

	// xlog.Debugf("Trying DIRECT connection to peer %s...", hop.ID[:maxIDChars])
	pc, err := pnhm.direct.connect(p2pHost)
	if err != nil {
		if hop.IsRelay {
			// xlog.Warnf("Unable to establish DIRECT connection with peer %s: %v", hop.ID, err)
			return nil, errors.Wrapf(err, "[%v] function pnhm.relay.connect()", errors.Trace())
		}
	} else {
		pc.DirectConnection = true
		return pc, nil
	}

	// xlog.Debugf("Trying RELAYED connection to peer %s...", hop.ID[:maxIDChars])
	pc, err = pnhm.relay.connect(p2pHost)
	if err != nil {
		// xlog.Warnf("Unable to establish RELAYED connection with peer %s: %v", hop.ID, err)
		return nil, errors.Wrapf(err, "[%v] function pnhm.relay.connect()", errors.Trace())
	}

	pc.DirectConnection = false

	return pc, nil
}
