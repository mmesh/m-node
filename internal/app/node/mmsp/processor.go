package mmsp

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-node/internal/app/node/mmsp/protos/nodemgmt"
	"mmesh.dev/m-node/internal/app/node/mmsp/protos/routing"
	"mmesh.dev/m-node/internal/app/node/mmsp/protos/workflow"
)

var RxQueue = make(chan *mmsp.Payload, 128)

func Processor(ctx context.Context, p *mmsp.Payload) {
	switch p.Type {
	case mmsp.PDUType_ROUTING:
		routing.Processor(ctx, p.RoutingPDU)
	case mmsp.PDUType_NODEMGMT:
		nodemgmt.Processor(ctx, p.NodeMgmtPDU)
	case mmsp.PDUType_WORKFLOW:
		workflow.Processor(ctx, p.WorkflowPDU)
	}
}
