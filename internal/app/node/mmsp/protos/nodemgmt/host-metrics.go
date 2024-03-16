package nodemgmt

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/kvstore/db/metricsdb"
)

func mmpHostMetricsRequest(pdu *mmsp.NodeMgmtPDU) error {
	if pdu.HostMetricsRequest == nil {
		return fmt.Errorf("null hostMetrisRequest")
	}
	req := pdu.HostMetricsRequest

	xlog.Debugf("[mmp] Received new host metrics request..")

	metricsdb.RequestQueue <- req

	return nil
}
