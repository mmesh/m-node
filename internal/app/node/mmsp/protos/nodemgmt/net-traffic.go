package nodemgmt

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/kvstore/db/netflowdb"
)

func mmpNetTrafficMetricsRequest(pdu *mmsp.NodeMgmtPDU) error {
	if pdu.NetTrafficMetricsRequest == nil {
		return fmt.Errorf("null netTrafficMetricsRequest")
	}
	req := pdu.NetTrafficMetricsRequest

	xlog.Debugf("[mmp] Received new traffic metrics request..")

	netflowdb.RequestQueue <- req

	return nil
}
