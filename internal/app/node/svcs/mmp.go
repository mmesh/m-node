package svcs

import (
	"context"
	"io"

	"github.com/spf13/viper"
	mmp_pb "mmesh.dev/api-go/grpc/network/mmp"
	"mmesh.dev/mmesh/internal/app/node/netp2p"
	"mmesh.dev/mmesh/internal/pkg/mmp"
	"mmesh.dev/mmesh/internal/pkg/runtime"
	"x6a.dev/pkg/xlog"
)

// mmpAgent method implementation of NxNetwork gRPC Service

func MMPAgent(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 2)

	go func() {
		//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancel()
		stream, err := w.NxNC.Cmd(context.Background())
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

				mmp.RecvCommandQueue <- payload
			}
			if err := stream.CloseSend(); err != nil {
				xlog.Errorf("Unable to close mmp stream: %v", err)
			}
			endCh <- struct{}{}
		}()

		go func() {
			for {
				select {
				case payload := <-mmp.SendCommandQueue:
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

		mmp.SendCommandQueue <- &mmp_pb.Payload{
			SrcID:       viper.GetString("mm.id"),
			PayloadType: mmp.PayloadTypeNodeInit,
			Node:        netp2p.GetNodeWithoutEndpoints(),
		}
	}()

	<-w.QuitChan

	endCh <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}
