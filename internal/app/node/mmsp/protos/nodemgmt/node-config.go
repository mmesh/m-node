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
	if pdu.NodeConfig == nil {
		return fmt.Errorf("null nodeConfig")
	}

	xlog.Infof("[mmp] Received new configuration..")

	if err := mnet.NewCfg(pdu.NodeConfig); err != nil {
		return errors.Wrapf(err, "[%v] function mnet.NewCfg()", errors.Trace())
	}

	return nil
}
