package portfwd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/mmp/stream/utils/cli"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func portFwdDisabled(payload *mmsp.Payload) {
	mmID := viper.GetString("mm.id")

	p := &mmsp.Payload{
		SrcID: mmID,
		DstID: payload.SrcID,
		Type: mmsp.PDUType_PORTFWD,
		PortFwdPDU: &mmsp.PortFwdPDU{
			Type: mmsp.PortFwdMsgType_PORTFWD_DISABLED,
		},
	}

	// p := &mmsp.Payload{
	// 	SrcID:       mmID,
	// 	DstID:       payload.SrcID,
	// 	PayloadType: mmsp.PayloadType_PORTFWD_DISABLED,
	// }

	queuing.TxControlQueue <- p
}

func portFwdUnavailable() {
	text := "Port forwarding unauthorized or disabled here"

	alert := fmt.Sprintf("\n%s%s%s", colors.Black("["), colors.Red(text), colors.Black("]"))
	fmt.Println(alert)

	cli.Disconnected()

	os.Exit(0)
}
