package routing

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/xlog"
)

func Processor(ctx context.Context, pdu *mmsp.RoutingPDU) {
	if pdu == nil {
		return
	}

	var err error

	switch pdu.Type {
	case mmsp.RoutingMsgType_ROUTING_STATUS:
		err = mmpRoutingStatus(ctx, pdu)
	}

	if err != nil {
		xlog.Errorf("[mmp] Unable to process mmp routingPDU (%s): %v",
			pdu.Type.String(), err)
	}
}
