package streaming

import "sync"

// const buffer_len = 1024
const BufferLen = 1048576

type IOSessionID string

type IOSessions struct {
	IO     map[IOSessionID]*IOPipes
	WaitCh map[IOSessionID]chan struct{}
	sync.RWMutex
}

func NewIOSessions() *IOSessions {
	return &IOSessions{
		IO:     make(map[IOSessionID]*IOPipes),
		WaitCh: make(map[IOSessionID]chan struct{}),
	}
}

func (ios *IOSessions) SetIOSession(sID string) {
	ios.Lock()
	defer ios.Unlock()

	ios.IO[IOSessionID(sID)] = NewIOPipes()
	ios.WaitCh[IOSessionID(sID)] = make(chan struct{})
}

func (ios *IOSessions) DeleteIOSession(sID string) {
	ios.Lock()
	defer ios.Unlock()

	if _, ok := ios.IO[IOSessionID(sID)]; ok {
		CloseIOPipes(ios.IO[IOSessionID(sID)])
		delete(ios.IO, IOSessionID(sID))
	}
	if _, ok := ios.WaitCh[IOSessionID(sID)]; ok {
		close(ios.WaitCh[IOSessionID(sID)])
		delete(ios.WaitCh, IOSessionID(sID))
	}
}

func (ios *IOSessions) GetIOSessionIO(sID string) *IOPipes {
	ios.Lock()
	defer ios.Unlock()

	io, ok := ios.IO[IOSessionID(sID)]
	if ok {
		return io
	}

	return nil
}

/*
func (ios *IOSessions) GetIOSessionWaitCh(sID string) *chan struct{} {
	ios.Lock()
	defer ios.Unlock()

	ch, ok := ios.WaitCh[IOSessionID(sID)]
	if ok {
		return &ch
	}

	return nil
}
*/
