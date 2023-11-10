package mmsp

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
)

func Preprocessor(ctx context.Context, p *mmsp.Payload) {
	switch p.Type {
	case mmsp.PDUType_ROUTING:
		if p.RoutingPDU == nil {
			return
		}

		switch p.RoutingPDU.Type {
		case mmsp.RoutingMsgType_ROUTING_STATUS:
			RxQueue <- p
			return
		}
	case mmsp.PDUType_NODEMGMT:
		if p.NodeMgmtPDU == nil {
			return
		}

		switch p.NodeMgmtPDU.Type {
		case mmsp.NodeMgmtMsgType_NODE_CONFIG:
			RxQueue <- p
			return
		case mmsp.NodeMgmtMsgType_NODE_METRICS_REQUEST:
			RxQueue <- p
			return
		}
	case mmsp.PDUType_WORKFLOW:
		if p.WorkflowPDU == nil {
			return
		}

		switch p.WorkflowPDU.Type {
		case mmsp.WorkflowMsgType_WORKFLOW_EXPEDITE:
			RxQueue <- p
			return
		case mmsp.WorkflowMsgType_WORKFLOW_SCHEDULE:
			RxQueue <- p
			return
		}
	case mmsp.PDUType_EVENT:
		if p.EventPDU == nil {
			return
		}
	}

	queuing.RxControlQueue <- p
}
