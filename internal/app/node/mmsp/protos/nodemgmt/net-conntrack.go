package nodemgmt

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/kvstore/db/ctlogdb"
	"mmesh.dev/m-node/internal/app/node/mnet/router/conntrack"
)

func mmpNetConntrackStateRequest(pdu *mmsp.NodeMgmtPDU) error {
	if pdu.NetCtStateRequest == nil {
		return fmt.Errorf("null netCtStateRequest")
	}
	req := pdu.NetCtStateRequest

	xlog.Debugf("[mmp] Received new conntrack state request..")

	conntrack.RequestQueue <- req

	return nil
}

func mmpNetConntrackLogRequest(pdu *mmsp.NodeMgmtPDU) error {
	if pdu.NetCtLogRequest == nil {
		return fmt.Errorf("null netCtLogRequest")
	}
	req := pdu.NetCtLogRequest

	xlog.Debugf("[mmp] Received new conntrack log request..")

	ctlogdb.RequestQueue <- req

	return nil
}
