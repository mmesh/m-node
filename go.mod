module mmesh.dev/m-node

go 1.15

replace mmesh.dev/m-api-go => ./modules/m-api-go

replace mmesh.dev/m-lib => ./modules/m-lib

require (
	github.com/c9s/goprocinfo v0.0.0-20190309065803-0b2ad9ac246b
	github.com/google/gopacket v1.1.19
	github.com/ipfs/go-datastore v0.4.5
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/libp2p/go-libp2p v0.13.0
	github.com/libp2p/go-libp2p-circuit v0.4.0
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/libp2p/go-libp2p-kad-dht v0.11.1
	github.com/libp2p/go-libp2p-swarm v0.4.0
	github.com/lorenzosaino/go-sysctl v0.1.0
	github.com/miekg/dns v1.1.31
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/onsi/ginkgo v1.14.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.0
	github.com/songgao/water v0.0.0-20190725173103-fd331bda3f4b
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.7.1
	github.com/vishvananda/netlink v1.1.0
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	google.golang.org/grpc v1.34.0
	k8s.io/api v0.20.1
	k8s.io/apimachinery v0.20.1
	k8s.io/client-go v0.20.1
	mmesh.dev/m-api-go v0.0.0-00010101000000-000000000000
	mmesh.dev/m-lib v0.0.0-00010101000000-000000000000
	x6a.dev/pkg v0.0.0-20200513233835-8ad67ebc7e71
)
