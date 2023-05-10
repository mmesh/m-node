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
	case mmsp.NodeMgmtMsgType_NODE_METRICS_REQUEST:
		err = mmpNodeMetricsRequest(pdu)
	}

	if err != nil {
		xlog.Errorf("[mmp] Unable to process mmp nodeMgmtPDU (%s): %v",
			pdu.Type.String(), err)
	}
}
