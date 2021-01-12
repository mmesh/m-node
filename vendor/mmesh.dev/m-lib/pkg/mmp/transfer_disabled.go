package mmp

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"x6a.dev/pkg/colors"
)

func transferDisabled(payload *mmsp.Payload) {
	mmID := viper.GetString("mm.id")

	p := &mmsp.Payload{
		SrcID:       mmID,
		DstID:       payload.SrcID,
		PayloadType: PayloadTypeTransferDisabled,
	}

	SendCommandQueue <- p
}

func transferUnavailable() {
	text := "File transfers unauthorized or disabled here"

	alert := fmt.Sprintf("\n%s%s%s", colors.Black("["), colors.Red(text), colors.Black("]"))
	fmt.Println(alert)

	Disconnected()

	os.Exit(0)
}
