package mmp

import "sync"

//const buffer_len = 1024
const buffer_len = 1048576

type shellSessionID string

type shellSession struct {
	io     map[shellSessionID]*ioPipes
	waitCh map[shellSessionID]chan struct{}
	sync.RWMutex
}

func newShellSession() *shellSession {
	return &shellSession{
		io:     make(map[shellSessionID]*ioPipes),
		waitCh: make(map[shellSessionID]chan struct{}),
	}
}

func (ss *shellSession) setShellSession(sID string) {
	ss.Lock()
	ss.io[shellSessionID(sID)] = newIOPipes()
	ss.waitCh[shellSessionID(sID)] = make(chan struct{})
	ss.Unlock()
}

func (ss *shellSession) deleteShellSession(sID string) {
	ss.Lock()
	if _, ok := ss.io[shellSessionID(sID)]; ok {
		closeIOPipes(ss.io[shellSessionID(sID)])
		delete(ss.io, shellSessionID(sID))
	}
	if _, ok := ss.waitCh[shellSessionID(sID)]; ok {
		close(ss.waitCh[shellSessionID(sID)])
		delete(ss.waitCh, shellSessionID(sID))
	}
	ss.Unlock()
}

func (ss *shellSession) getShellSessionIO(sID string) *ioPipes {
	ss.Lock()
	defer ss.Unlock()

	io, ok := ss.io[shellSessionID(sID)]
	if ok {
		return io
	}

	return nil
}

func (ss *shellSession) getShellSessionWaitCh(sID string) *chan struct{} {
	ss.Lock()
	defer ss.Unlock()

	ch, ok := ss.waitCh[shellSessionID(sID)]
	if ok {
		return &ch
	}

	return nil
}
