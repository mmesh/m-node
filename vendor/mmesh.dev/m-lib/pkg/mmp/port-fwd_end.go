package mmp

import (
	"context"
	"sync"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/portFwd"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/logging"
)

func ClosePortFwd(srcPort, dstPort uint32, proto, srcNodeID, dstNodeID string, wg *sync.WaitGroup) {
	logging.Debugf("Closing port-forward session: srcID %s, dstID %s", srcNodeID, dstNodeID)

	mmID := viper.GetString("mm.id")

	ipProto := uint32(ipnet.IPProtocol(proto))
	pfLinkID := portFwdLinkID(ipProto, srcNodeID, srcPort, dstNodeID, dstPort)

	p1 := newPortFwdEnd(mmID, srcNodeID, srcNodeID, dstNodeID, pfLinkID)
	TxControlQueue <- p1

	p2 := newPortFwdEnd(mmID, dstNodeID, srcNodeID, dstNodeID, pfLinkID)
	TxControlQueue <- p2

	logging.Debug("Cleaning port-forward connections...")
	time.Sleep(time.Second)

	wg.Done()
}

func portFwdClose(p *mmsp.Payload) {
	mmID := viper.GetString("mm.id")
	srcNodeID := p.PortFwd.Link.SrcNodeID
	dstNodeID := p.PortFwd.Link.DstNodeID
	lID := p.PortFwd.Link.ID

	p1 := newPortFwdEnd(mmID, p.PortFwd.Link.SrcNodeID, srcNodeID, dstNodeID, lID)
	TxControlQueue <- p1

	p2 := newPortFwdEnd(mmID, p.PortFwd.Link.DstNodeID, srcNodeID, dstNodeID, lID)
	TxControlQueue <- p2
}

func newPortFwdEnd(srcID, dstID, srcNodeID, dstNodeID, pfLinkID string) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       dstID,
		PayloadType: mmsp.PayloadType_PORTFWD_END,
		PortFwd: &portFwd.PortFwd{
			Link: &portFwd.Link{
				ID:        pfLinkID,
				SrcNodeID: srcNodeID,
				DstNodeID: dstNodeID,
			},
		},
	}
}

func portFwdEnd(ctx context.Context, p *mmsp.Payload) error {
	mmID := viper.GetString("mm.id")
	srcNodeID := p.PortFwd.Link.SrcNodeID
	dstNodeID := p.PortFwd.Link.DstNodeID
	lID := p.PortFwd.Link.ID

	if mmID == srcNodeID {
		lpfs.deletePortFwdSession(lID)
	}
	if mmID == dstNodeID {
		rpfs.deletePortFwdSession(lID)
	}

	logging.Debugf("Port-Forward connection disconnected: srcID %s, dstID %s", srcNodeID, dstNodeID)

	return nil
}
