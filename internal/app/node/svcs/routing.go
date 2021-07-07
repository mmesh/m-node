package svcs

import (
	"context"
	"io"
	"time"

	"mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-node/internal/app/node/netp2p"
	"x6a.dev/pkg/xlog"
)

// RoutingAgent runs routing engine
func RoutingAgent(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 2)

	go func() {
		//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancel()
		stream, err := w.NxNC.Routing(context.Background())
		if err != nil {
			xlog.Errorf("Unable to get routing stream from mmesh controller: %v", err)
			networkErrorEventsQueue <- struct{}{}
			return
		}

		go func() {
			for {
				rtResp, err := stream.Recv()
				if err == io.EOF {
					xlog.Warnf("Ended (io.EOF) routing stream: %v", err)
					networkErrorEventsQueue <- struct{}{}
					break
				}
				if err != nil {
					xlog.Warnf("Unable to receive rtResponse payload: %v", err)
					networkErrorEventsQueue <- struct{}{}
					break
				}
				//xlog.Tracef("Routing table (VRF %s) updated", rtResp.RT.VRFID)

				netp2p.UpdateRoutingTable(rtResp.RT)
			}
			if err := stream.CloseSend(); err != nil {
				xlog.Errorf("Unable to close mmp stream: %v", err)
			}
			endCh <- struct{}{}
		}()

		go func() {
			rtrq := netp2p.GetRTRQueue()
			for {
				select {
				case <-rtrq:
					rtReq := &routing.RTRequest{
						Node:        netp2p.GetNode(),
						Connections: netp2p.GetLinkStatusConnections(),
					}
					if err := stream.Send(rtReq); err != nil {
						xlog.Warnf("Unable to send rtRequest payload: %v", err)
						networkErrorEventsQueue <- struct{}{}
						if err := stream.CloseSend(); err != nil {
							xlog.Errorf("Unable to close mmp stream: %v", err)
						}
						break
					}
				case <-endCh:
					xlog.Debug("Closing rtRequest send stream")
					return
				}
			}
		}()
	}()

	go rtCtl()

	<-w.QuitChan

	endCh <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}

var rtCtlRun bool

func rtCtl() {
	if !rtCtlRun {
		rtCtlRun = true
		rtrq := netp2p.GetRTRQueue()
		for {
			rtrq <- struct{}{}
			time.Sleep(20 * time.Second)
		}
	}
}
