package start

import (
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/update"
	"mmesh.dev/m-node/internal/app/node/ops"
	"mmesh.dev/m-node/internal/app/node/svcs"
)

const (
	errorEventsHandler = iota
	// networkErrorEventsHandler
	mmDispatcher
	mmProcessor
	dnsAgent
	metricsAgent
	mmpControlAgent
	routingAgent
	cronAgent
	atdAgent
	k8sConnector
	// proxy64gc
	federationMonitor
	updateAgent
	// bgpAgent
)

func initWrkrs(nxnc rpc.NetworkAPIClient) {
	runtime.RegisterWrkr(
		errorEventsHandler,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmErrorEventsHandler"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, runtime.ErrorEventsHandler),
	)
	// runtime.RegisterWrkr(
	// 	networkErrorEventsHandler,
	// 	runtime.SetWrkrOpt(runtime.WrkrOptName, "mmNetworkErrorEventsHandler"),
	// 	runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.NetworkErrorEventsHandler),
	// )
	runtime.RegisterWrkr(
		mmDispatcher,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmDispatcher"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, mmp.Dispatcher),
	)
	runtime.RegisterWrkr(
		mmProcessor,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmProcessor"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.MMPProcessor),
	)
	runtime.RegisterWrkr(
		dnsAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmDNSAgent"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.DNSAgent),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		metricsAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmMetricsAgent"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.MetricsAgent),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		mmpControlAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmpControlAgent"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.MMPControl),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		routingAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmRoutingAgent"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.RoutingAgent),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		cronAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmCron"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, ops.Cron),
	)
	runtime.RegisterWrkr(
		atdAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmAtd"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, ops.Atd),
	)
	runtime.RegisterWrkr(
		k8sConnector,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmKubernetesGateway"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.KubernetesConnector),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	// runtime.RegisterWrkr(
	// 	proxy64gc,
	// 	runtime.SetWrkrOpt(runtime.WrkrOptName, "mmProxy64GC"),
	// 	runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.Proxy64GC),
	// )
	runtime.RegisterWrkr(
		federationMonitor,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmFederationMonitor"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.FederationMonitor),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		updateAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmUpdateAgent"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, update.UpdateAgent),
	)
	// runtime.RegisterWrkr(
	// 	bgpAgent,
	// 	runtime.SetWrkrOpt(runtime.WrkrOptName, "mmBGPAgent"),
	// 	runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.BGPAgent),
	// )
}
