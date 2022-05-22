package svcs

import (
	"context"
	"io"
	"os"
	"time"

	"mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet"
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
			xlog.Errorf("Unable to get routing stream from controller: %v", err)
			mnet.LocalNode().Connection().Watcher() <- struct{}{}
			return
		}

		go func() {
			for {
				rtResp, err := stream.Recv()
				if err == io.EOF {
					// xlog.Warnf("Ended (io.EOF) routing stream: %v", err)
					mnet.LocalNode().Connection().Watcher() <- struct{}{}
					break
				}
				if err != nil {
					// xlog.Warnf("Unable to receive rtResponse: %v", err)
					mnet.LocalNode().Connection().Watcher() <- struct{}{}
					break
				}

				if rtResp.RT.Disabled {
					xlog.Alert("Service is DISABLED")
					xlog.Alert("Please contact mmesh customer service urgently")
					// break
					os.Exit(1)
				}

				mnet.LocalNode().Router().UpdateRoutingTable(rtResp.RT)
			}
			if err := stream.CloseSend(); err != nil {
				xlog.Errorf("Unable to close routing stream: %v", err)
			}
			endCh <- struct{}{}
			// xlog.Warn("Closing rtResponse recv stream")
		}()

		go func() {
			for {
				select {
				case <-rtCtlQueue:
					rtReq := &routing.RTRequest{
						Node:        mnet.LocalNode().NetworkNode(),
						Connections: mnet.LocalNode().Router().GetLinkStatusConnections(),
					}
					if err := stream.Send(rtReq); err != nil {
						// xlog.Warnf("Unable to send rtRequest: %v", err)
						mnet.LocalNode().Connection().Watcher() <- struct{}{}
						if err := stream.CloseSend(); err != nil {
							xlog.Errorf("Unable to close routing stream: %v", err)
						}
						return
					}
				case <-endCh:
					// xlog.Warn("Closing rtRequest send stream")
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
			rtrq := mnet.LocalNode().Router().RtrQueue()
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
