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

var rtCtlQueue = make(chan struct{})

// RoutingAgent runs routing engine
func RoutingAgent(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 2)

	go func() {
		// ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		// defer cancel()
		stream, err := w.NxNC.Routing(context.TODO())
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
					xlog.Warnf("Unable to receive rtResponse: %v", err)
					networkErrorEventsQueue <- struct{}{}
					break
				}

				netp2p.UpdateRoutingTable(rtResp.RT)
			}
			if err := stream.CloseSend(); err != nil {
				xlog.Errorf("Unable to close routing stream: %v", err)
			}
			endCh <- struct{}{}
			xlog.Warn("Closing rtResponse recv stream")
		}()

		go func() {
			for {
				select {
				case <-rtCtlQueue:
					rtReq := &routing.RTRequest{
						Node:        netp2p.GetNode(),
						Connections: netp2p.GetLinkStatusConnections(),
					}
					if err := stream.Send(rtReq); err != nil {
						xlog.Warnf("Unable to send rtRequest: %v", err)
						networkErrorEventsQueue <- struct{}{}
						if err := stream.CloseSend(); err != nil {
							xlog.Errorf("Unable to close routing stream: %v", err)
						}
						return
					}
				case <-endCh:
					xlog.Warn("Closing rtRequest send stream")
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

		go func() {
			for {
				rtCtlQueue <- struct{}{}
				time.Sleep(20 * time.Second)
			}
		}()

		go func() {
			rtrqCtlRun := false
			rtrq := netp2p.GetRTRQueue()
			for {
				<-rtrq
				if !rtrqCtlRun {
					rtrqCtlRun = true

					go func() {
						rtCtlQueue <- struct{}{}
						time.Sleep(5 * time.Second)
						rtrqCtlRun = false
					}()
				}
			}
		}()
	}
}
