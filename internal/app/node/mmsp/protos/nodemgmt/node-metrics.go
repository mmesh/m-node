package nodemgmt

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/tss"
)

func mmpNodeMetricsRequest(pdu *mmsp.NodeMgmtPDU) error {
	if pdu.NodeMetricsRequest == nil {
		return fmt.Errorf("null nodeMetrisRequest")
	}
	nmreq := pdu.NodeMetricsRequest

	xlog.Infof("[mmp] Received new metrics request..")

	tss.RequestQueue <- nmreq

	return nil
}
