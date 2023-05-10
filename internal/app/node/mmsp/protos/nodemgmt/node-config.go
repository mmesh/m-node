package nodemgmt

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet"
)

func mmpNodeConfig(ctx context.Context, pdu *mmsp.NodeMgmtPDU) error {
	if pdu.NodeCfg == nil {
		return fmt.Errorf("null nodeCfg")
	}
	ncfg := pdu.NodeCfg

	xlog.Infof("[mmp] Received new configuration..")
	if err := mnet.LocalNode().NewCfg(ncfg); err != nil {
		return errors.Wrapf(err, "[%v] function mnet.LocalNode().NewCfg()", errors.Trace())
	}

	return nil
}
