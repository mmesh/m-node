package mmp

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
)

var RxControlQueue = make(chan *mmsp.Payload, 128)
var TxControlQueue = make(chan *mmsp.Payload, 128)

var ControllerQueue = make(chan *mmsp.Payload, 128)
var NodeQueue = make(chan *mmsp.Payload, 128)

func Dispatcher(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	close := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case payload := <-RxControlQueue:
				xlog.Debugf("Received pdu on controlQueue from %s", payload.SrcID)
				Processor(context.TODO(), payload)
			case <-close:
				xlog.Debug("Closing mmp dispacher")
				return
			}
		}
	}()

	<-w.QuitChan

	close <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}

func Processor(ctx context.Context, p *mmsp.Payload) error {
	var err error

	switch p.PayloadType {
	case mmsp.PayloadType_NODE_INIT:
		ControllerQueue <- p
	case mmsp.PayloadType_NODE_KEEPALIVE:
		ControllerQueue <- p

	case mmsp.PayloadType_COMMAND_SHELL_EXEC:
		go shellExec(ctx, p)
	case mmsp.PayloadType_COMMAND_SHELL_EXIT:
		err = shellExit(ctx, p)
	case mmsp.PayloadType_COMMAND_SHELL_INPUT:
		err = shellReadInput(ctx, p)
	case mmsp.PayloadType_COMMAND_SHELL_OUTPUT:
		err = shellReadOutput(ctx, p)
	case mmsp.PayloadType_COMMAND_SHELL_DISABLED:
		shellUnavailable()

	case mmsp.PayloadType_TRANSFER_INIT:
		err = transferInit(ctx, p)
	case mmsp.PayloadType_TRANSFER_DATA:
		err = transferDataRx(ctx, p)
	case mmsp.PayloadType_TRANSFER_ACK:
		err = transferAckRx(ctx, p)
	case mmsp.PayloadType_TRANSFER_DISABLED:
		transferUnavailable()

	case mmsp.PayloadType_PORTFWD_LISTEN:
		go portFwdListen(ctx, p)
	case mmsp.PayloadType_PORTFWD_DIAL:
		go portFwdDial(ctx, p)
	case mmsp.PayloadType_PORTFWD_DIALACK:
		err = portFwdDialAck(ctx, p)
	case mmsp.PayloadType_PORTFWD_END:
		err = portFwdEnd(ctx, p)
	case mmsp.PayloadType_PORTFWD_DATA:
		err = portFwdReadData(ctx, p)
	case mmsp.PayloadType_PORTFWD_DISABLED:
		portFwdUnavailable()

	case mmsp.PayloadType_WORKFLOW_EXPEDITE:
		NodeQueue <- p
	case mmsp.PayloadType_WORKFLOW_SCHEDULE:
		NodeQueue <- p
	case mmsp.PayloadType_WORKFLOW_RESPONSE:
		ControllerQueue <- p

	case mmsp.PayloadType_ALERT_EVENT:
		ControllerQueue <- p

	case mmsp.PayloadType_CHAT_SESSION_INIT:
		// Init runs on controllers
		ControllerQueue <- p
		// go ChatSession().Init(p)
	case mmsp.PayloadType_CHAT_SESSION_END:
		// End runs on controllers
		ControllerQueue <- p
		// go ChatSession().End(p)
	case mmsp.PayloadType_CHAT_SESSION_EXIT:
		// Exit runs on user cli
		ChatSession().Exit()
	case mmsp.PayloadType_CHAT_SESSION_INPUT:
		// Init runs on controllers
		err = ChatSession().Input(p)
	case mmsp.PayloadType_CHAT_SESSION_OUTPUT:
		// Exit runs on user cli
		err = ChatSession().Output(p)
	case mmsp.PayloadType_CHAT_SESSION_DISABLED:
		// Exit runs on user cli
		ChatSession().Disabled()
	}

	if err != nil {
		logging.Debugf("Unable to process mmp payload: %v", err)
		return err
	}

	return nil
}
