package connection

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"mmesh.dev/m-api-go/grpc/network/mmnp/natProbe"
	"mmesh.dev/m-api-go/grpc/network/rpc"
	"x6a.dev/pkg/errors"
)

func agentNATProbe(nxnc rpc.NxNetworkClient, natp *natProbe.NATProbe) (*natProbe.NATProbe, error) {
	var wg sync.WaitGroup

	ctx := context.Background()
	waitc := make(chan struct{})

	wg.Add(1)
	go probeServer(strconv.Itoa(int(natp.Port)), &wg, waitc)

	time.Sleep(1 * time.Second)

	natp, err := nxnc.NATProbe(ctx, natp)
	if err != nil {
		waitc <- struct{}{}
		wg.Wait()
		return natp, errors.Wrapf(err, "[%v] function nxnc.NATProbe()", errors.Trace())
	}

	waitc <- struct{}{}
	wg.Wait()

	return natp, nil
}

func probeServer(port string, wg *sync.WaitGroup, waitc chan struct{}) {
	smux := http.NewServeMux()

	smux.HandleFunc("/probe", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Ok\n")
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: smux,
	}

	go func() {
		server.ListenAndServe()
	}()

	<-waitc
	server.Shutdown(context.Background())
	wg.Done()
}
