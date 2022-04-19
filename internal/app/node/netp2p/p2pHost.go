package netp2p

import (
	"context"
	"crypto/rand"

	// mrand "math/rand"

	// ds "github.com/ipfs/go-datastore"
	// dsync "github.com/ipfs/go-datastore/sync"
	"github.com/libp2p/go-libp2p"

	// connmgr "github.com/libp2p/go-libp2p-connmgr"

	circuit "github.com/libp2p/go-libp2p-circuit"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	// "github.com/libp2p/go-libp2p-core/routing"
	// dht "github.com/libp2p/go-libp2p-kad-dht"

	// relayv1 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv1/relay"
	// relayv2 "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
	// "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"

	// libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	// secio "github.com/libp2p/go-libp2p-secio"
	// libp2ptls "github.com/libp2p/go-libp2p-tls"
	// rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	"mmesh.dev/m-lib/pkg/errors"
)

const (
	p2pHostTypeBasicHost = iota
	p2pHostTypeHiddenHost
	p2pHostTypeRelayHost
)

func newP2PHost(port int32, p2pHostType int) (host.Host, error) {
	r := rand.Reader
	// r := mrand.New(mrand.NewSource(int64(port)))

	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)", errors.Trace())
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
		// libp2p.Transport(libp2pquic.NewTransport),
		libp2p.DefaultTransports,
		libp2p.DefaultMuxers,
		// support TLS connections
		// libp2p.Security(libp2ptls.ID, libp2ptls.New),
		// support secio connections
		// libp2p.Security(secio.ID, secio.New),
		libp2p.DefaultSecurity,
		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		// libp2p.ConnectionManager(cm),
		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),
		// Let this host use the DHT to find other hosts (required by autoRelay)
		// libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
		// 	return dht.New(context.TODO(), h)
		// }),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		// libp2p.EnableAutoRelay(),
	}

	switch p2pHostType {
	case p2pHostTypeHiddenHost:
		opts = append(opts,
			// libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
			libp2p.ListenAddrs(),
			// Enable the use of relays
			libp2p.EnableRelay(),
			libp2p.ForceReachabilityPrivate(),
		)
	case p2pHostTypeBasicHost:
		opts = append(opts,
			libp2p.ListenAddrs(getLocalMAddrs(port)...),
			// Enable the use of relays
			libp2p.EnableRelay(),
		)
	case p2pHostTypeRelayHost:
		opts = append(opts,
			libp2p.ListenAddrs(getLocalMAddrs(port)...),
			libp2p.EnableRelay(circuit.OptHop),
			// libp2p.EnableRelayService(), // circuitv2
			libp2p.ForceReachabilityPublic(),
		)
	}

	host, err := libp2p.New(context.TODO(), opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function libp2p.New()", errors.Trace())
	}

	/*
		if p2pHostType == p2pHostTypeRelayHost {
			// if _, err := relayv1.NewRelay(
			// 	host,
			// 	relayv1.WithResources(relayv1.DefaultResources()),
			// ); err != nil {
			// 	return nil, errors.Wrapf(err, "[%v] function relayv1.NewRelay()", errors.Trace())
			// }
			// if _, err := relayv2.New(
			// 	host,
			// 	relayv2.WithResources(relayv2.DefaultResources()),
			// ); err != nil {
			// 	return nil, errors.Wrapf(err, "[%v] function relayv2.New()", errors.Trace())
			// }
		}
	*/

	// Construct a datastore (needed by the DHT). This is just a simple, in-memory thread-safe datastore.
	// dstore := dsync.MutexWrap(ds.NewMapDatastore())

	// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	// mmDHT := dht.NewDHT(ctx, host, dstore)

	// Make the routed host
	// routedHost := rhost.Wrap(host, mmDHT)
	// host = routedHost

	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.

	// xlog.Trace("Bootstrapping DHT...")
	// if err = mmDHT.Bootstrap(ctx); err != nil {
	// 	return nil, errors.Wrapf(err, "[%v] function nxDHT.Bootstrap()", errors.Trace())
	// }

	return host, nil
}
