package rib

import (
	"time"

	"mmesh.dev/m-lib/pkg/xlog"
)

func (r *ribData) wrkr(msgQueue <-chan []byte) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case m := <-msgQueue:
			if err := r.decoder(string(m)); err != nil {
				xlog.Errorf("Unable to decode RIB data: %v", err)
			}
		case d := <-r.rxQueue:
			r.processor(d)
		case <-ticker.C:
			r.cleanup()
		case <-r.closeCh:
			return
		}
	}
}
