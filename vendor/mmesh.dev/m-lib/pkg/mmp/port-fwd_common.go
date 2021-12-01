package mmp

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/portFwd"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp/streaming"
)

func portFwdLinkID(proto, srcID, srcPort, dstID, dstPort interface{}) string {
	pfLinkID := []byte(fmt.Sprintf("%v:%v:%v:%v:%v", proto, srcID, srcPort, dstID, dstPort))

	return base64.URLEncoding.EncodeToString(pfLinkID)
}

func newPortFwdData(srcID, dstID, cID string, p *mmsp.Payload) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       dstID,
		RequesterID: p.RequesterID,
		Interactive: p.Interactive,
		PayloadType: mmsp.PayloadType_PORTFWD_DATA,
		PortFwd: &portFwd.PortFwd{
			Link: &portFwd.Link{
				ID:           p.PortFwd.Link.ID,
				Proto:        p.PortFwd.Link.Proto,
				SrcNodeID:    p.PortFwd.Link.SrcNodeID,
				SrcPort:      p.PortFwd.Link.SrcPort,
				DstNodeID:    p.PortFwd.Link.DstNodeID,
				DstPort:      p.PortFwd.Link.DstPort,
				ConnectionID: cID,
			},
			Data: nil,
		},
	}
}

func portFwdWriteData(srcID, dstID, cID string, payload *mmsp.Payload, outrp *io.PipeReader) {
	for {
		p := newPortFwdData(srcID, dstID, cID, payload)
		buffer := make([]byte, streaming.BufferLen)
		n, err := outrp.Read(buffer)
		if err == io.EOF {
			//logging.Tracef("OUT Pipe Reader EOF: %v", err)
			outrp.Close()
			break
		}
		if err != nil {
			//logging.Tracef("OUT Pipe Reader: %v", err)
			outrp.Close()
			break
		}

		logging.Tracef("Sending %v bytes of data to %s, connectionID %s", n, dstID, cID)

		p.PortFwd.Data = buffer[0:n]
		TxControlQueue <- p
	}
}

func portFwdReadData(ctx context.Context, payload *mmsp.Payload) error {
	var iop *streaming.IOPipes

	if len(payload.PortFwd.Data) == 0 {
		logging.Trace("len(payload.PortFwd.Data) == 0")
		return nil
	}

	mmID := viper.GetString("mm.id")
	srcID := payload.SrcID
	srcNodeID := payload.PortFwd.Link.SrcNodeID
	dstNodeID := payload.PortFwd.Link.DstNodeID
	lID := payload.PortFwd.Link.ID
	cID := payload.PortFwd.Link.ConnectionID

	// logging.Tracef("Processing data from ConnectionID %s", cID)

	if mmID == srcNodeID {
		if lpfs == nil {
			return errors.Errorf("port-fwd connectionID %s not initialized", cID)
		}
		iop = lpfs.getPortFwdConnIO(lID, cID)
	}
	if mmID == dstNodeID {
		if rpfs == nil {
			return errors.Errorf("port-fwd connectionID %s not initialized", cID)
		}
		iop = rpfs.getPortFwdConnIO(lID, cID)
	}
	if iop == nil {
		return errors.Errorf("port-fwd input writer pipe not found for node %s and connection %s", srcID, cID)
	}
	if iop.In.WP == nil {
		return errors.Errorf("port-fwd input writer pipe not found for node %s and connection %s", srcID, cID)
	}

	n, err := iop.In.WP.Write(payload.PortFwd.Data)
	if err != nil {
		// logging.Errorf("Unable to read inbound data from node %s: %v", srcID, err)
		return errors.Wrapf(err, "[%v] function iop.in.wp.Write(payload.PortFwd.Data)", errors.Trace())
	}
	logging.Tracef("Received %d bytes from node %s, connectionID %s", n, srcID, cID)

	return nil
}
