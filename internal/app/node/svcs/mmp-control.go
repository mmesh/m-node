package svcs

import (
	"context"
	"io"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-node/internal/app/node/netp2p"
	"x6a.dev/pkg/xlog"
)

// Control method implementation of NetworkAPI gRPC Service
func MMPControl(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 2)

	go func() {
		//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancel()
		stream, err := w.NxNC.Control(context.TODO())
		if err != nil {
			xlog.Errorf("Unable to get mmp stream from mmesh controller: %v", err)
			networkErrorEventsQueue <- struct{}{}
			return
		}

		go func() {
			for {
				payload, err := stream.Recv()
				if err == io.EOF {
					xlog.Warnf("Ended (io.EOF) mmp stream: %v", err)
					networkErrorEventsQueue <- struct{}{}
					break
				}
				if err != nil {
					xlog.Warnf("Unable to receive mmp payload: %v", err)
					networkErrorEventsQueue <- struct{}{}
					break
				}
				// xlog.Tracef("Received mmp payload from %s", payload.SrcID)

				mmp.RxControlQueue <- payload
			}
			if err := stream.CloseSend(); err != nil {
				xlog.Errorf("Unable to close mmp stream: %v", err)
			}
			endCh <- struct{}{}
		}()

		go func() {
			for {
				select {
				case payload := <-mmp.TxControlQueue:
					if err := stream.Send(payload); err != nil {
						xlog.Warnf("Unable to send mmp payload: %v", err)
						networkErrorEventsQueue <- struct{}{}
						if err := stream.CloseSend(); err != nil {
							xlog.Errorf("Unable to close mmp stream: %v", err)
						}
						break
					}
				case <-endCh:
					xlog.Debug("Closing mmp send stream")
					return
				}
			}
		}()

		mmp.TxControlQueue <- &mmsp.Payload{
			SrcID:       viper.GetString("mm.id"),
			PayloadType: mmsp.PayloadType_NODE_INIT,
			Node:        netp2p.GetNodeWithoutEndpoints(),
		}
	}()

	go mmpCtl()

	<-w.QuitChan

	endCh <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}

var mmpCtlRun bool

func mmpCtl() {
	if !mmpCtlRun {
		mmpCtlRun = true
		for {
			mmp.TxControlQueue <- &mmsp.Payload{
				SrcID:       viper.GetString("mm.id"),
				PayloadType: mmsp.PayloadType_NODE_KEEPALIVE,
			}
			time.Sleep(60 * time.Second)
		}
	}
}
