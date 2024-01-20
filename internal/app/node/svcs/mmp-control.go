package svcs

import (
	"context"
	"io"
	"time"

	"github.com/spf13/viper"
	mmsp_pb "mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mmsp"
	"mmesh.dev/m-node/internal/app/node/mnet"
)

// Control method implementation of NetworkAPI gRPC Service
func NetworkControl(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 2)

	go func() {
		stream, err := w.NxNC.Control(context.Background())
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

				// if !serviceEnabled {
				// 	continue
				// }

				mmsp.Preprocessor(context.TODO(), payload)
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
				case payload := <-mmsp.RxQueue:
					xlog.Debug("[mmp] Received mmp payload on queue")
					go mmsp.Processor(context.TODO(), payload)

				case payload := <-queuing.TxControlQueue:
					if err := stream.Send(payload); err != nil {
						// xlog.Warnf("[mmp] Unable to send mmp payload: %v", err)

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

		queuing.TxControlQueue <- &mmsp_pb.Payload{
			SrcID: viper.GetString("mm.id"),
			Type:  mmsp_pb.PDUType_NODEMGMT,
			NodeMgmtPDU: &mmsp_pb.NodeMgmtPDU{
				Type:    mmsp_pb.NodeMgmtMsgType_NODE_INIT,
				NodeReq: mnet.LocalNode().NodeReq(),
			},
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
	mmID := viper.GetString("mm.id")

	if !mnet.LocalNode().Node().Cfg.DisableNetworking || mnet.LocalNode().Router() != nil {
		return
	}

	if !mmpCtlRun {
		mmpCtlRun = true
		for {
			queuing.TxControlQueue <- &mmsp_pb.Payload{
				SrcID: mmID,
				Type:  mmsp_pb.PDUType_SESSION,
				SessionPDU: &mmsp_pb.SessionPDU{
					Type:      mmsp_pb.SessionMsgType_SESSION_KEEPALIVE,
					SessionID: mmID,
				},
			}
			time.Sleep(30 * time.Second)
		}
	}
}
