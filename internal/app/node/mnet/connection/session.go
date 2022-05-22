package connection

import (
	"context"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmnp/natprobe"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (c *connection) newSession() error {
	// hiddenNode := viper.GetBool("agent.hidden")

	c.externalIPv4 = viper.GetString("agent.externalIPv4") // could be ""

	natp := &natprobe.NATProbe{
		Port:         int32(viper.GetInt("agent.port")),
		ExternalIPv4: c.externalIPv4,
	}

	natp, err := c.natProbe(natp)
	if err != nil {
		xlog.Alertf("Unable to initialize agent: %v", errors.Cause(err))
		os.Exit(1)
	}

	// only ipv4 for now
	if strings.Contains(natp.ExternalIPv4, ":") {
		return nil
	}

	// configured externalIPv4 overrides controller's detected externalIPv4
	if len(c.externalIPv4) == 0 {
		c.externalIPv4 = natp.ExternalIPv4
	}

	// if hiddenNode {
	// 	c.externalIPv4 = ""
	// }

	return nil
}

func (c *connection) natProbe(natp *natprobe.NATProbe) (*natprobe.NATProbe, error) {
	var wg sync.WaitGroup

	ctx := context.TODO()
	waitc := make(chan struct{})

	wg.Add(1)
	go probeServer(strconv.Itoa(int(natp.Port)), &wg, waitc)

	time.Sleep(1 * time.Second)

	natp, err := c.nxnc.NATProbe(ctx, natp)
	if err != nil {
		waitc <- struct{}{}
		wg.Wait()
		return natp, errors.Wrapf(err, "[%v] function c.nxnc.NATProbe()", errors.Trace())
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
	server.Shutdown(context.TODO())

	wg.Done()
}
