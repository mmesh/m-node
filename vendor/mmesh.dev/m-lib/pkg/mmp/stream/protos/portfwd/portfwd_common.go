package portfwd

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/stream"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp/stream/utils/streaming"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
)

func portFwdLinkID(proto, srcID, srcPort, dstID, dstPort interface{}) string {
	pfLinkID := []byte(fmt.Sprintf("%v:%v:%v:%v:%v", proto, srcID, srcPort, dstID, dstPort))

	return base64.URLEncoding.EncodeToString(pfLinkID)
}

func newPortFwdData(srcID, dstID, cID string, p *mmsp.Payload) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       dstID,
		Type: mmsp.PDUType_PORTFWD,
		PortFwdPDU: &mmsp.PortFwdPDU{
			Metadata: &mmsp.StreamMetadata{
				RequesterID: p.PortFwdPDU.Metadata.RequesterID,
				Interactive: p.PortFwdPDU.Metadata.Interactive,
			},
			Type: mmsp.PortFwdMsgType_PORTFWD_DATA,
			PortFwd: &stream.PortFwd{
				Link: &stream.Link{
					ID:           p.PortFwdPDU.PortFwd.Link.ID,
					Proto:        p.PortFwdPDU.PortFwd.Link.Proto,
					SrcNodeID:    p.PortFwdPDU.PortFwd.Link.SrcNodeID,
					SrcPort:      p.PortFwdPDU.PortFwd.Link.SrcPort,
					DstNodeID:    p.PortFwdPDU.PortFwd.Link.DstNodeID,
					DstPort:      p.PortFwdPDU.PortFwd.Link.DstPort,
					ConnectionID: cID,
				},
				Data: nil,
			},
		},
	}

	// return &mmsp.Payload{
	// 	SrcID:       srcID,
	// 	DstID:       dstID,
	// 	RequesterID: p.RequesterID,
	// 	Interactive: p.Interactive,
	// 	PayloadType: mmsp.PayloadType_PORTFWD_DATA,
	// 	PortFwd: &portFwd.PortFwd{
	// 		Link: &portFwd.Link{
	// 			ID:           p.PortFwd.Link.ID,
	// 			Proto:        p.PortFwd.Link.Proto,
	// 			SrcNodeID:    p.PortFwd.Link.SrcNodeID,
	// 			SrcPort:      p.PortFwd.Link.SrcPort,
	// 			DstNodeID:    p.PortFwd.Link.DstNodeID,
	// 			DstPort:      p.PortFwd.Link.DstPort,
	// 			ConnectionID: cID,
	// 		},
	// 		Data: nil,
	// 	},
	// }
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

		p.PortFwdPDU.PortFwd.Data = buffer[0:n]
		queuing.TxControlQueue <- p
	}
}

func portFwdReadData(ctx context.Context, payload *mmsp.Payload) error {
	var iop *streaming.IOPipes

	if len(payload.PortFwdPDU.PortFwd.Data) == 0 {
		logging.Trace("len(payload.PortFwdPDU.PortFwd.Data) == 0")
		return nil
	}

	mmID := viper.GetString("mm.id")
	srcID := payload.SrcID
	srcNodeID := payload.PortFwdPDU.PortFwd.Link.SrcNodeID
	dstNodeID := payload.PortFwdPDU.PortFwd.Link.DstNodeID
	lID := payload.PortFwdPDU.PortFwd.Link.ID
	cID := payload.PortFwdPDU.PortFwd.Link.ConnectionID

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

	n, err := iop.In.WP.Write(payload.PortFwdPDU.PortFwd.Data)
	if err != nil {
		// logging.Errorf("Unable to read inbound data from node %s: %v", srcID, err)
		return errors.Wrapf(err, "[%v] function iop.in.wp.Write(payload.PortFwd.Data)", errors.Trace())
	}
	logging.Tracef("Received %d bytes from node %s, connectionID %s", n, srcID, cID)

	return nil
}
