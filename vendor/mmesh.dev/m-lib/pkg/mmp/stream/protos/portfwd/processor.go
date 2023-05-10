package portfwd

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
)

func Processor(ctx context.Context, p *mmsp.Payload) error {
	if p.PortFwdPDU == nil {
		return fmt.Errorf("null pdu")
	}

	var err error

	switch p.PortFwdPDU.Type {
	case mmsp.PortFwdMsgType_PORTFWD_LISTEN:
		go portFwdListen(ctx, p)
	case mmsp.PortFwdMsgType_PORTFWD_DIAL:
		go portFwdDial(ctx, p)
	case mmsp.PortFwdMsgType_PORTFWD_DIALACK:
		err = portFwdDialAck(ctx, p)
	case mmsp.PortFwdMsgType_PORTFWD_END:
		err = portFwdEnd(ctx, p)
	case mmsp.PortFwdMsgType_PORTFWD_DATA:
		err = portFwdReadData(ctx, p)
	case mmsp.PortFwdMsgType_PORTFWD_DISABLED:
		portFwdUnavailable()
	}

	if err != nil {
		logging.Errorf("Unable to process mmp portFwdPDU payload: %v", errors.Cause(err))
		return err
	}

	return nil
}
