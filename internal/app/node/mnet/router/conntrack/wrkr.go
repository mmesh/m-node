package conntrack

import "time"

func (ctm *ctMap) wrkr() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctm.cleanup()
		case <-ctm.closeCh:
			return
		}
	}
}
