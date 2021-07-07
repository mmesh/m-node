package mmp

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"x6a.dev/pkg/colors"
)

func shellExecDisabled(payload *mmsp.Payload) {
	mmID := viper.GetString("mm.id")

	p := &mmsp.Payload{
		SrcID:       mmID,
		DstID:       payload.SrcID,
		PayloadType: mmsp.PayloadType_COMMAND_SHELL_DISABLED,
	}

	TxControlQueue <- p
}

func shellUnavailable() {
	text := "Shell exec unauthorized or disabled here"

	alert := fmt.Sprintf("%s%s%s", colors.Black("["), colors.Red(text), colors.Black("]"))
	fmt.Println(alert)

	Disconnected()

	os.Exit(0)
}
