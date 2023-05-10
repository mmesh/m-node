package metrics

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/resources/events"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/xlog"
)

func newAlertEvent(evt *events.Event) {
	mmID := viper.GetString("mm.id")

	xlog.Debugf("[event] New event from srcID %s", mmID)

	queuing.TxControlQueue <- &mmsp.Payload{
		SrcID: mmID,
		Type:  mmsp.PDUType_EVENT,
		EventPDU: &mmsp.EventPDU{
			Event: evt,
		},
	}
}
