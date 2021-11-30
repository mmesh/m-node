package mmp

import (
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/alert"
	"mmesh.dev/m-lib/pkg/xlog"
)

func NewEvent(eventPayload *alert.EventPayload) {
	mmID := eventPayload.SourceID

	xlog.Debugf("[Alert] New event from srcID %s", mmID)

	p := &mmsp.Payload{
		SrcID:       mmID,
		RequesterID: mmID,
		PayloadType: mmsp.PayloadType_ALERT_EVENT,
		Event:       eventPayload,
	}

	TxControlQueue <- p
}
