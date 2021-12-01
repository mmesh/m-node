package netp2p

import (
	"context"
	"crypto/rand"

	//mrand "math/rand"

	ds "github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	"github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	// libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	// secio "github.com/libp2p/go-libp2p-secio"
	// libp2ptls "github.com/libp2p/go-libp2p-tls"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

const (
	p2pHostTypeBasicHost = iota
	p2pHostTypeHiddenHost
	p2pHostTypeRelayHost
)

func newP2PHost(port int32, p2pHostType int) (host.Host, error) {
	var opts []libp2p.Option

	r := rand.Reader
	//r := mrand.New(mrand.NewSource(int64(port)))

	// Creates a new RSA key pair for this host.
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)", errors.Trace())
	}

	// 0.0.0.0 will listen on any interface device.
	// sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "[%v] function multiaddr.NewMultiaddr()", errors.Trace())
	// }

	switch p2pHostType {
	case p2pHostTypeHiddenHost:
		// sourceMultiAddr, err := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/0")
		// if err != nil {
		// 	return nil, errors.Wrapf(err, "[%v] function multiaddr.NewMultiaddr()", errors.Trace())
		// }

		opts = []libp2p.Option{
			//libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
			libp2p.ListenAddrs(),
			//libp2p.ListenAddrs(sourceMultiAddr),
			libp2p.Identity(prvKey),
			// libp2p.EnableRelay(circuit.OptDiscovery),
			libp2p.EnableAutoRelay(),
			// libp2p.Transport(libp2pquic.NewTransport),
			libp2p.DefaultTransports,
			libp2p.DefaultMuxers,
			// support TLS connections
			// libp2p.Security(libp2ptls.ID, libp2ptls.New),
			// support secio connections
			// libp2p.Security(secio.ID, secio.New),
			libp2p.DefaultSecurity,
			libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
				return dht.New(context.TODO(), h)
			}),
			libp2p.NATPortMap(),
		}
	case p2pHostTypeBasicHost:
		opts = []libp2p.Option{
			libp2p.ListenAddrs(getLocalMAddrs(port)...),
			libp2p.Identity(prvKey),
			// libp2p.EnableRelay(circuit.OptDiscovery),
			libp2p.EnableAutoRelay(),
			// libp2p.Transport(libp2pquic.NewTransport),
			libp2p.DefaultTransports,
			libp2p.DefaultMuxers,
			// support TLS connections
			// libp2p.Security(libp2ptls.ID, libp2ptls.New),
			// support secio connections
			// libp2p.Security(secio.ID, secio.New),
			libp2p.DefaultSecurity,
			libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
				return dht.New(context.Background(), h)
			}),
			libp2p.NATPortMap(),
		}
	case p2pHostTypeRelayHost:
		opts = []libp2p.Option{
			libp2p.ListenAddrs(getLocalMAddrs(port)...),
			libp2p.Identity(prvKey),
			// libp2p.EnableRelay(circuit.OptDiscovery),
			libp2p.EnableRelay(circuit.OptHop),
			libp2p.EnableAutoRelay(),
			// libp2p.Transport(libp2pquic.NewTransport),
			libp2p.DefaultTransports,
			libp2p.DefaultMuxers,
			// support TLS connections
			// libp2p.Security(libp2ptls.ID, libp2ptls.New),
			// support secio connections
			// libp2p.Security(secio.ID, secio.New),
			libp2p.DefaultSecurity,
			libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
				return dht.New(context.Background(), h)
			}),
			libp2p.NATPortMap(),
		}
	}

	ctx := context.TODO()

	host, err := libp2p.New(ctx, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function libp2p.New()", errors.Trace())
	}

	// Construct a datastore (needed by the DHT). This is just a simple, in-memory thread-safe datastore.
	dstore := dsync.MutexWrap(ds.NewMapDatastore())

	// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	nxDHT := dht.NewDHT(ctx, host, dstore)

	// Make the routed host
	routedHost := rhost.Wrap(host, nxDHT)
	host = routedHost

	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.
	//logger.Debug("Bootstrapping the DHT")
	xlog.Trace("Bootstrapping DHT...")
	if err = nxDHT.Bootstrap(ctx); err != nil {
		return nil, errors.Wrapf(err, "[%v] function nxDHT.Bootstrap(ctx)", errors.Trace())
	}

	return host, nil
}
