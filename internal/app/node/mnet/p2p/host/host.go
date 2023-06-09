package host

import (
	"crypto/rand"
	"fmt"
	"time"

	// mrand "math/rand"

	"github.com/libp2p/go-libp2p"

	// connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"

	// circuit "github.com/libp2p/go-libp2p-circuit"
	// relayv1 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv1/relay"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
	// "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"

	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	tcp "github.com/libp2p/go-libp2p/p2p/transport/tcp"

	// secio "github.com/libp2p/go-libp2p-secio"
	// libp2ptls "github.com/libp2p/go-libp2p-tls"
	"mmesh.dev/m-lib/pkg/errors"
)

type P2PHostType int

const (
	P2PHostTypeHiddenHost P2PHostType = iota
	P2PHostTypeBasicHost
	P2PHostTypeRelayHost
)

func New(hostType P2PHostType, port int) (host.Host, error) {
	r := rand.Reader
	// r := mrand.New(mrand.NewSource(int64(port)))

	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function crypto.GenerateKeyPairWithReader()", errors.Trace())
	}

	// cm, err := connmgr.NewConnManager(
	// 	512,                                    // Lowwater
	// 	768,                                    // HighWater,
	// 	connmgr.WithGracePeriod(2*time.Minute), // GracePeriod
	// )
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "[%v] function connmgr.NewConnManager()", errors.Trace())
	// }

	opts := []libp2p.Option{
		libp2p.Identity(prvKey),
		// libp2p.DefaultTransports,
		libp2p.ChainOptions(
			libp2p.Transport(quic.NewTransport),
			libp2p.Transport(tcp.NewTCPTransport),
		),
		libp2p.DefaultMuxers,
		// support TLS connections
		// libp2p.Security(libp2ptls.ID, libp2ptls.New),
		// support secio connections
		// libp2p.Security(secio.ID, secio.New),
		libp2p.DefaultSecurity,
		// Enable the use of relays
		libp2p.EnableRelay(),
		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		// libp2p.ConnectionManager(cm),
		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),
		// Experimental EnableHolePunching enables NAT traversal
		// by enabling NATT'd peers to both initiate and respond
		// to hole punching attempts to create direct/NAT-traversed
		// connections with other peers. (default: disabled)
		libp2p.EnableHolePunching(),
		// DisableMetrics configures libp2p to disable prometheus metrics
		libp2p.DisableMetrics(),
	}

	maddrs := []string{
		fmt.Sprintf("/ip4/0.0.0.0/udp/%d/quic-v1", port), // UDP endpoint for the QUIC transport
		fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port),         // regular TCP connections
	}

	switch hostType {
	case P2PHostTypeHiddenHost:
		opts = append(opts,
			// libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
			libp2p.ListenAddrs(),
			libp2p.ForceReachabilityPrivate(),
		)
	case P2PHostTypeBasicHost:
		opts = append(opts,
			libp2p.ListenAddrStrings(maddrs...),
		)
	case P2PHostTypeRelayHost:
		opts = append(opts,
			libp2p.ListenAddrStrings(maddrs...),
			libp2p.EnableRelayService(
				relay.WithResources(relay.Resources{
					Limit: &relay.RelayLimit{
						Duration: 4 * time.Hour, // default: 2 * time.Minute
						// Duration: 1 * time.Minute, // default: 2 * time.Minute
						Data: 1 << 30, // 1GB  // default: 1 << 17, // 128K
					},

					ReservationTTL: time.Hour,

					MaxReservations: 128, // default: 128
					MaxCircuits:     16,
					BufferSize:      2048,

					MaxReservationsPerPeer: 4,
					MaxReservationsPerIP:   8,
					MaxReservationsPerASN:  32,
				}),
				// relay.WithResources(relay.DefaultResources()),
				// relay.WithInfiniteLimits(),
			), // circuitv2
			libp2p.ForceReachabilityPublic(),
		)
	}

	h, err := libp2p.New(opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function libp2p.New()", errors.Trace())
	}

	// if hostType == P2PHostTypeRelayHost {
	// 	if _, err := relay.New(
	// 		h,
	// 		relay.WithResources(relay.DefaultResources()),
	// 		relay.WithInfiniteLimits(),
	// 	); err != nil {
	// 		return nil, errors.Wrapf(err, "[%v] function relay.New()", errors.Trace())
	// 	}
	// }

	return h, nil
}
