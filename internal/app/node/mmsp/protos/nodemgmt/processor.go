package nodemgmt

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/xlog"
)

func Processor(ctx context.Context, pdu *mmsp.NodeMgmtPDU) {
	if pdu == nil {
		return
	}

	var err error

	switch pdu.Type {
	case mmsp.NodeMgmtMsgType_NODE_CONFIG:
		err = mmpNodeConfig(ctx, pdu)
	case mmsp.NodeMgmtMsgType_NODE_HOST_METRICS_REQUEST:
		err = mmpHostMetricsRequest(pdu)
	case mmsp.NodeMgmtMsgType_NODE_NET_CONNTRACK_STATE_REQUEST:
		err = mmpNetConntrackStateRequest(pdu)
	case mmsp.NodeMgmtMsgType_NODE_NET_CONNTRACK_LOG_REQUEST:
		err = mmpNetConntrackLogRequest(pdu)
	case mmsp.NodeMgmtMsgType_NODE_NET_TRAFFIC_METRICS_REQUEST:
		err = mmpNetTrafficMetricsRequest(pdu)
	case mmsp.NodeMgmtMsgType_NODE_HOST_SECURITY_REQUEST:
		err = mmpHostSecurityReportRequest(pdu)
	}

	if err != nil {
		xlog.Errorf("[mmp] Unable to process mmp nodeMgmtPDU (%s): %v",
			pdu.Type.String(), err)
	}
}
