package utils

import (
	"os"
	"os/signal"
	"sync"
)

// WaitForCtrlC waits until an os.Interrupt signal is sent (ctrl + c)
func WaitForCtrlC() {
	var wg sync.WaitGroup
	var signalCh chan os.Signal

	signalCh = make(chan os.Signal, 1)

	wg.Add(1)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		<-signalCh
		wg.Done()
	}()

	wg.Wait()
}
