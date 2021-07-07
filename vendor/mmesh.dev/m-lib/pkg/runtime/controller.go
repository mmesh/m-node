package runtime

import (
	"context"
	"sync"
	"time"

	"mmesh.dev/m-api-go/grpc/rpc"
	"x6a.dev/pkg/xlog"
)

const (
	WrkrOptName = iota
	WrkrOptJobQueue
	WrkrOptStartFunc
	WrkrOptExtFunc
	WrkrOptNxNetworkClient
)

type WrkrOpt struct {
	Key   int
	Value interface{}
}

type Wrkr struct {
	id        int
	Name      string
	JobQueue  interface{}
	QuitChan  chan struct{}
	Running   bool
	WG        *sync.WaitGroup
	Ctx       context.Context
	startFunc func(*Wrkr)
	extFunc   func()
	NxNC      rpc.NetworkAPIClient
}

type nxWrkr interface {
	start()
	stop()
}

var wList = make(map[int][]*WrkrOpt)
var wrkrs = make(map[int]*Wrkr)
var wwg sync.WaitGroup

func RegisterWrkr(wID int, wOpts ...*WrkrOpt) {
	for _, wOpt := range wOpts {
		wList[wID] = append(wList[wID], wOpt)
	}
}

func SetWrkrOpt(k int, v interface{}) *WrkrOpt {
	return &WrkrOpt{
		Key:   k,
		Value: v,
	}
}

func (w *Wrkr) setWrkrOptions(wOpts ...*WrkrOpt) {
	for _, wOpt := range wOpts {
		switch wOpt.Key {
		case WrkrOptName:
			w.Name = wOpt.Value.(string)
		case WrkrOptJobQueue:
			w.JobQueue = wOpt.Value.(chan Job)
		case WrkrOptStartFunc:
			w.startFunc = wOpt.Value.(func(*Wrkr))
		case WrkrOptExtFunc:
			w.extFunc = wOpt.Value.(func())
		case WrkrOptNxNetworkClient:
			w.NxNC = wOpt.Value.(rpc.NetworkAPIClient)
		}
	}
}

// newWrkr creates a new worker
func newWrkr(ctx context.Context, wID int, wOpts ...*WrkrOpt) *Wrkr {
	w := new(Wrkr)
	w.Ctx = ctx
	w.id = wID
	w.setWrkrOptions(wOpts...)
	w.QuitChan = make(chan struct{})

	return w
}

func (w *Wrkr) start() {
	go func() {
		xlog.Debugf("Starting worker %s...", w.Name)
		for w.Running {
			time.Sleep(time.Second)
		}
		w.startFunc(w)
	}()
}

func (w *Wrkr) stop() {
	xlog.Debugf("Stopping worker %s...", w.Name)
	w.QuitChan <- struct{}{}
}

func StartWrkrs() {
	for wID, wOpts := range wList {
		ctx := context.Background()
		w := newWrkr(ctx, wID, wOpts...)
		wwg.Add(1)
		w.WG = &wwg
		w.start()
		wrkrs[wID] = w
	}
}

func StopWrkrs(wg *sync.WaitGroup) {
	for wID, w := range wrkrs {
		w.stop()
		delete(wrkrs, wID)
	}
	wwg.Wait()
	wg.Done()
}

func NetworkWrkrReconnect(newNxNC rpc.NetworkAPIClient) {
	for _, w := range wrkrs {
		if w.NxNC != nil {
			w.stop()
			w.NxNC = newNxNC
			wwg.Add(1)
			w.WG = &wwg
			w.start()
		}
	}
}
