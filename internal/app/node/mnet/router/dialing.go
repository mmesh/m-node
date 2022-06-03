package router

func (r *router) setDial(dstAddr string) bool {
	r.dialing.Lock()
	defer r.dialing.Unlock()

	if _, ok := r.dialing.addr[dstAddr]; ok {
		return false
	}

	r.dialing.addr[dstAddr] = struct{}{}

	return true
}

func (r *router) unsetDial(dstAddr string) {
	r.dialing.Lock()
	defer r.dialing.Unlock()

	delete(r.dialing.addr, dstAddr)
}
