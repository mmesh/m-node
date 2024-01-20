package routing

import (
	"context"
	"fmt"
	"os"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/xlog"
)

var ServiceEnabled bool = true
var disabledRetries int

func mmpRoutingStatus(ctx context.Context, pdu *mmsp.RoutingPDU) error {
	if pdu.Status == nil {
		return fmt.Errorf("null status")
	}
	s := pdu.Status

	if s.Disabled {
		xlog.Alert("Service is DISABLED.")
		xlog.Alert("Please contact mmesh customer service urgently.")
	}
	// } else if s.OverLimit {
	// 	xlog.Alert("Account over tier limits. Service is DISABLED.")
	// 	xlog.Alert("If you are on the Free Plan, make sure you")
	// 	xlog.Alert("are not exceeding its limits. If not, please")
	// 	xlog.Alert("contact mmesh customer service urgently.")
	// }

	// if s.Disabled || s.OverLimit {
	if s.Disabled {
		ServiceEnabled = false

		disabledRetries++
		if disabledRetries > 10 {
			os.Exit(1)
		}
		return nil
	} else {
		ServiceEnabled = true
		disabledRetries = 0
	}

	return nil
}
