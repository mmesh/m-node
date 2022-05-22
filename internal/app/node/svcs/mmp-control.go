package svcs

import (
	"context"
	"io"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet"
)

// Control method implementation of NetworkAPI gRPC Service
func MMPControl(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 2)

	go func() {
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// defer cancel()
		stream, err := w.NxNC.Control(context.TODO())
		if err != nil {
			xlog.Errorf("Unable to get mmp stream from controller: %v", err)
			mnet.LocalNode().Connection().Watcher() <- struct{}{}
			return
		}

		go func() {
			for {
				payload, err := stream.Recv()
				if err == io.EOF {
					// xlog.Warnf("Ended (io.EOF) mmp stream: %v", err)
					mnet.LocalNode().Connection().Watcher() <- struct{}{}
					break
				}
				if err != nil {
					// xlog.Warnf("Unable to receive mmp payload: %v", err)
					mnet.LocalNode().Connection().Watcher() <- struct{}{}
					break
				}

				mmp.RxControlQueue <- payload
			}
			if err := stream.CloseSend(); err != nil {
				xlog.Errorf("Unable to close mmp stream: %v", err)
			}
			endCh <- struct{}{}
			// xlog.Warn("Closing mmp recv stream")
		}()

		go func() {
			for {
				select {
				case payload := <-mmp.TxControlQueue:
					if err := stream.Send(payload); err != nil {
						// xlog.Warnf("Unable to send mmp payload: %v", err)
						mnet.LocalNode().Connection().Watcher() <- struct{}{}
						if err := stream.CloseSend(); err != nil {
							xlog.Errorf("Unable to close mmp stream: %v", err)
						}
						return
					}
				case <-endCh:
					// xlog.Warn("Closing mmp send stream")
					return
				}
			}
		}()

		mmp.TxControlQueue <- &mmsp.Payload{
			SrcID:       viper.GetString("mm.id"),
			PayloadType: mmsp.PayloadType_NODE_INIT,
			Node:        mnet.LocalNode().NetworkNodeWithoutEndpoints(),
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
