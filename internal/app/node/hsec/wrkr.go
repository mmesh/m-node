package hsec

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/resources/nstore/hsecdb"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/runtime"
	"mmesh.dev/m-lib/pkg/xlog"
)

var RequestQueue = make(chan *hsecdb.HostSecurityReportRequest, 128)

func Scanner(w *runtime.Wrkr) {
	xlog.Infof("Started worker %s", w.Name)
	w.Running = true

	endCh := make(chan struct{}, 2)

	go func() {
		securityScannerCtl := make(chan struct{}, 1)
		go func() {
			time.Sleep(30 * time.Second)
			securityScannerCtl <- struct{}{}
		}()

		mmID := viper.GetString("mm.id")

		ticker := time.NewTicker(24 * 3600 * time.Second) // 24 hours
		defer ticker.Stop()

		for {
			select {
			case <-securityScannerCtl:
				if err := scan(); err != nil {
					xlog.Warnf("[host-security] Unable to complete security scan: %v", errors.Cause(err))
					continue
				}

			case <-ticker.C:
				securityScannerCtl <- struct{}{}

			case r := <-RequestQueue:
				hsr, err := readReportFile()
				if err != nil {
					xlog.Warnf("[host-security] Unable to get host security report: %v", errors.Cause(err))
				}

				hsrr := query(r, hsr) // hsr can be nil

				queuing.TxControlQueue <- &mmsp.Payload{
					SrcID:           mmID,
					DstControllerID: r.Request.ControllerID,
					Type:            mmsp.PDUType_NODEMGMT,
					NodeMgmtPDU: &mmsp.NodeMgmtPDU{
						Type:               mmsp.NodeMgmtMsgType_NODE_HOST_SECURITY_RESPONSE,
						HsecReportResponse: hsrr,
					},
				}

			case <-endCh:
				// xlog.Warn("[host-security] Closing security scanner")
				return
			}
		}
	}()

	<-w.QuitChan

	endCh <- struct{}{}

	w.WG.Done()
	w.Running = false
	xlog.Infof("Stopped worker %s", w.Name)
}
