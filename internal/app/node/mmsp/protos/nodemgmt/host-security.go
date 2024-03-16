package nodemgmt

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/hsec"
)

func mmpHostSecurityReportRequest(pdu *mmsp.NodeMgmtPDU) error {
	if pdu.HsecReportRequest == nil {
		return fmt.Errorf("null hsecReportRequest")
	}
	req := pdu.HsecReportRequest

	xlog.Debugf("[mmp] Received new host security report request..")

	hsec.RequestQueue <- req

	return nil
}
