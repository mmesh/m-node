package router

import "net/netip"

func (r *router) setDial(dstAddr netip.Addr) bool {
	r.dialing.Lock()
	defer r.dialing.Unlock()

	if _, ok := r.dialing.addr[dstAddr]; ok {
		return false
	}

	r.dialing.addr[dstAddr] = struct{}{}

	return true
}

func (r *router) unsetDial(dstAddr netip.Addr) {
	r.dialing.Lock()
	defer r.dialing.Unlock()

	delete(r.dialing.addr, dstAddr)
}
