package mmp

import (
	"sync"
)

type pfConnectionID string
type pfLinkID string

type portFwdConnection struct {
	io        map[pfConnectionID]*ioPipes
	dialAckCh map[pfConnectionID]chan struct{}
	//sync.RWMutex
}

type portFwdSession struct {
	conns map[pfLinkID]*portFwdConnection
	sync.RWMutex
}

func newPortFwdConnection() *portFwdConnection {
	return &portFwdConnection{
		io:        make(map[pfConnectionID]*ioPipes),
		dialAckCh: make(map[pfConnectionID]chan struct{}),
	}
}

func newPortFwdSession() *portFwdSession {
	return &portFwdSession{
		conns: make(map[pfLinkID]*portFwdConnection),
	}
}

func (pfs *portFwdSession) setPortFwdSession(lID string) {
	pfs.Lock()
	pfs.conns[pfLinkID(lID)] = newPortFwdConnection()
	pfs.Unlock()
}

func (pfs *portFwdSession) deletePortFwdSession(lID string) {
	pfs.Lock()
	if pfl, ok := pfs.conns[pfLinkID(lID)]; ok {
		for cID := range pfl.io {
			closeIOPipes(pfl.io[pfConnectionID(cID)])
			delete(pfl.io, pfConnectionID(cID))
		}
		for cID := range pfl.dialAckCh {
			close(pfl.dialAckCh[pfConnectionID(cID)])
			delete(pfl.dialAckCh, pfConnectionID(cID))
		}
		delete(pfs.conns, pfLinkID(lID))
	}
	pfs.Unlock()
}

func (pfs *portFwdSession) setPortFwdConnection(lID, cID string) {
	pfs.Lock()
	if _, ok := pfs.conns[pfLinkID(lID)]; !ok {
		pfs.Unlock()
		pfs.setPortFwdSession(lID)
		pfs.Lock()
	}
	pfs.conns[pfLinkID(lID)].io[pfConnectionID(cID)] = newIOPipes()
	pfs.conns[pfLinkID(lID)].dialAckCh[pfConnectionID(cID)] = make(chan struct{})
	pfs.Unlock()
}

func (pfs *portFwdSession) deletePortFwdConnection(lID, cID string) {
	pfs.Lock()
	if pfl, ok1 := pfs.conns[pfLinkID(lID)]; ok1 {
		if ioPipes, ok2 := pfl.io[pfConnectionID(cID)]; ok2 {
			closeIOPipes(ioPipes)
			delete(pfl.io, pfConnectionID(cID))
		}
		if dialAckCh, ok2 := pfl.dialAckCh[pfConnectionID(cID)]; ok2 {
			close(dialAckCh)
			delete(pfl.dialAckCh, pfConnectionID(cID))
		}
	}
	pfs.Unlock()
}

func (pfs *portFwdSession) getPortFwdConnIO(lID, cID string) *ioPipes {
	pfs.Lock()
	defer pfs.Unlock()

	if pfl, ok1 := pfs.conns[pfLinkID(lID)]; ok1 {
		io, ok2 := pfl.io[pfConnectionID(cID)]
		if ok2 {
			return io
		}
	}
	return nil
}

func (pfs *portFwdSession) getPortFwdConnDialAckCh(lID, cID string) *chan struct{} {
	pfs.Lock()
	defer pfs.Unlock()

	if pfl, ok1 := pfs.conns[pfLinkID(lID)]; ok1 {
		ch, ok2 := pfl.dialAckCh[pfConnectionID(cID)]
		if ok2 {
			return &ch
		}
	}
	return nil
}
