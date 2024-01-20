package hstat

import (
	"net/netip"
)

func (hs *hstats) updateNet(addr netip.Addr, tx, rx uint64) {
	hs.Lock()
	defer hs.Unlock()

	// if _, ok := hs.byAddr[addr]; !ok {
	// 	hs.byAddr[addr] = &netTraffic{
	// 		rx: &netdb.TrafficCounter{},
	// 		tx: &netdb.TrafficCounter{},
	// 	}
	// }

	hs.netTraffic.tx.Bytes += tx
	hs.netTraffic.rx.Bytes += rx

	// hs.byAddr[addr].tx.Bytes = +tx
	// hs.byAddr[addr].rx.Bytes = +rx

	if tx > 0 {
		hs.netTraffic.tx.Packets++
		// hs.byAddr[addr].tx.Packets++
	}
	if rx > 0 {
		hs.netTraffic.rx.Packets++
		// hs.byAddr[addr].rx.Packets++
	}
}
