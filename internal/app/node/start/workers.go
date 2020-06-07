package start

import (
	"mmesh.dev/mmesh/internal/api/grpc/network/rpc"
	"mmesh.dev/mmesh/internal/app/node/ae/dispatcher"
	"mmesh.dev/mmesh/internal/app/node/svcs"
	"mmesh.dev/mmesh/internal/pkg/mmp"
	"mmesh.dev/mmesh/internal/pkg/runtime"
	"mmesh.dev/mmesh/internal/pkg/update"
)

const (
	errorEventsHandler = iota
	networkErrorEventsHandler
	mmProcessor
	mmWorkflowDispatcher
	dnsAgent
	metricsAgent
	mmpAgent
	routingAgent
	cronAgent
	atdAgent
	k8sConnector
	proxy64gc
	federationMonitor
	updateAgent
	//bgpAgent
)

func initWrkrs(nxnc rpc.NxNetworkClient) {
	runtime.RegisterWrkr(
		errorEventsHandler,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmErrorEventsHandler"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, runtime.ErrorEventsHandler),
	)
	runtime.RegisterWrkr(
		networkErrorEventsHandler,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmNetworkErrorEventsHandler"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.NetworkErrorEventsHandler),
	)
	runtime.RegisterWrkr(
		mmProcessor,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmProcessor"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, mmp.CmdEventsHandler),
	)
	runtime.RegisterWrkr(
		mmWorkflowDispatcher,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmWorkflowDispatcher"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, dispatcher.WorkflowEventHandler),
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
		mmpAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmpAgent"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.MMPAgent),
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
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, dispatcher.Cron),
	)
	runtime.RegisterWrkr(
		atdAgent,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmAtd"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, dispatcher.Atd),
	)
	runtime.RegisterWrkr(
		k8sConnector,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmKubernetesGateway"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.KubernetesConnector),
		runtime.SetWrkrOpt(runtime.WrkrOptNxNetworkClient, nxnc),
	)
	runtime.RegisterWrkr(
		proxy64gc,
		runtime.SetWrkrOpt(runtime.WrkrOptName, "mmProxy64GC"),
		runtime.SetWrkrOpt(runtime.WrkrOptStartFunc, svcs.Proxy64GC),
	)
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
