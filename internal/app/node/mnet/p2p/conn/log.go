package conn

import (
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p/core/network"
	"mmesh.dev/m-lib/pkg/xlog"
)

func Log(c network.Conn) {
	connType := "DIRECT"
	if c.Stat().Transient {
		connType = "INDIRECT"
	}

	xlog.Info("----------------------------------------------")
	xlog.Infof("New %s Connection: %s", strings.ToUpper(c.Stat().Direction.String()), c.ID())
	xlog.Infof("Connection Type: %s", connType)
	xlog.Infof("Local Peer: %s", c.LocalPeer().ShortString())
	xlog.Infof("Remote Peer: %s", c.RemotePeer().ShortString())

	if len(fmt.Sprintf("%v", c.ConnState().StreamMultiplexer)) > 0 {
		xlog.Infof("Multiplexer: %v", c.ConnState().StreamMultiplexer)
	}

	if len(fmt.Sprintf("%v", c.ConnState().Security)) > 0 {
		xlog.Infof("Security: %v", c.ConnState().Security)
	}

	xlog.Infof("Transport: %v", c.ConnState().Transport)
	xlog.Infof("Local MultiAddr: %v", c.LocalMultiaddr())
	xlog.Infof("Remote MultiAddr: %v", c.RemoteMultiaddr())
	xlog.Info("----------------------------------------------")
}
