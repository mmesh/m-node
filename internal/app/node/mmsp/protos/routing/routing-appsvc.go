package routing

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-node/internal/app/node/mnet"
)

func mmpRoutingAppSvcConfig(ctx context.Context, pdu *mmsp.RoutingPDU) error {
	if pdu.AppSvcConfig == nil {
		return fmt.Errorf("null appSvcConfig")
	}

	ascfg := pdu.AppSvcConfig

	switch ascfg.Operation {
	case mmsp.AppSvcConfigOperation_APPSVC_SET:
		mnet.LocalNode().Router().RIB().AddNodeAppSvc(ascfg.AppSvc)
	case mmsp.AppSvcConfigOperation_APPSVC_UNSET:
		mnet.LocalNode().Router().RIB().RemoveNodeAppSvc(ascfg.AppSvc.AppSvcID)
	}

	if mnet.LocalNode().Node().Cfg.DisableNetworking || mnet.LocalNode().Router() == nil {
		return nil
	}

	if !ServiceEnabled {
		return nil
	}

	mmID := viper.GetString("mm.id")

	mnet.LocalNode().SendAppSvcLSAs(mmID)

	return nil
}
