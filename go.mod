module mmesh.dev/m-node

go 1.15

replace mmesh.dev/m-api-go => ./modules/m-api-go

replace mmesh.dev/m-lib => ./modules/m-lib

require (
	github.com/c9s/goprocinfo v0.0.0-20190309065803-0b2ad9ac246b
	github.com/google/gopacket v1.1.19
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/ipfs/go-datastore v0.0.5
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/libp2p/go-libp2p v0.3.0
	github.com/libp2p/go-libp2p-circuit v0.1.1
	github.com/libp2p/go-libp2p-core v0.2.1
	github.com/libp2p/go-libp2p-kad-dht v0.2.0
	github.com/libp2p/go-libp2p-swarm v0.2.0
	github.com/lorenzosaino/go-sysctl v0.1.0
	github.com/miekg/dns v1.1.27
	github.com/multiformats/go-multiaddr v0.0.4
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.0
	github.com/songgao/water v0.0.0-20190725173103-fd331bda3f4b
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.7.1
	github.com/vishvananda/netlink v1.1.0
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	k8s.io/api v0.20.1
	k8s.io/apimachinery v0.20.1
	k8s.io/client-go v11.0.0+incompatible
	mmesh.dev/m-api-go v0.0.0-00010101000000-000000000000
	mmesh.dev/m-lib v0.0.0-00010101000000-000000000000
	x6a.dev/pkg v0.0.0-20200513233835-8ad67ebc7e71
)

replace k8s.io/client-go => k8s.io/client-go v0.20.1
