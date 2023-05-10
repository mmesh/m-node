package portfwd

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/google/gopacket/layers"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/stream"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp/stream/utils/auth"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
)

var rpfs = newPortFwdSession()

func newPortFwdDialAck(p *mmsp.Payload) *mmsp.Payload {
	// port-fwd dialAck must be sent by DstNodeID
	logging.Tracef("Sending DialAck: ConnectionID %s", p.PortFwdPDU.PortFwd.Link.ConnectionID)

	return &mmsp.Payload{
		SrcID: p.PortFwdPDU.PortFwd.Link.DstNodeID,
		DstID: p.PortFwdPDU.PortFwd.Link.SrcNodeID,
		Type:  mmsp.PDUType_PORTFWD,
		PortFwdPDU: &mmsp.PortFwdPDU{
			Metadata: &mmsp.StreamMetadata{
				RequesterID: p.PortFwdPDU.Metadata.RequesterID,
				Interactive: p.PortFwdPDU.Metadata.Interactive,
			},
			Type: mmsp.PortFwdMsgType_PORTFWD_DIALACK,
			PortFwd: &stream.PortFwd{
				Link: &stream.Link{
					ID:           p.PortFwdPDU.PortFwd.Link.ID,
					Proto:        p.PortFwdPDU.PortFwd.Link.Proto,
					SrcNodeID:    p.PortFwdPDU.PortFwd.Link.SrcNodeID,
					SrcPort:      p.PortFwdPDU.PortFwd.Link.SrcPort,
					DstNodeID:    p.PortFwdPDU.PortFwd.Link.DstNodeID,
					DstPort:      p.PortFwdPDU.PortFwd.Link.DstPort,
					ConnectionID: p.PortFwdPDU.PortFwd.Link.ConnectionID,
				},
				Data: nil,
			},
		},
	}

	// return &mmsp.Payload{
	// 	SrcID:       p.PortFwd.Link.DstNodeID,
	// 	DstID:       p.PortFwd.Link.SrcNodeID,
	// 	RequesterID: p.RequesterID,
	// 	Interactive: p.Interactive,
	// 	PayloadType: mmsp.PayloadType_PORTFWD_DIALACK,
	// 	PortFwd: &portFwd.PortFwd{
	// 		Link: &portFwd.Link{
	// 			ID:           p.PortFwd.Link.ID,
	// 			Proto:        p.PortFwd.Link.Proto,
	// 			SrcNodeID:    p.PortFwd.Link.SrcNodeID,
	// 			SrcPort:      p.PortFwd.Link.SrcPort,
	// 			DstNodeID:    p.PortFwd.Link.DstNodeID,
	// 			DstPort:      p.PortFwd.Link.DstPort,
	// 			ConnectionID: p.PortFwd.Link.ConnectionID,
	// 		},
	// 		Data: nil,
	// 	},
	// }
}

func portFwdConnDial(proto, dstPort uint32) (net.Conn, error) {
	var c net.Conn
	var err error

	ipv4Addr := net.JoinHostPort("127.0.0.1", fmt.Sprintf("%d", dstPort))

	switch proto {
	case uint32(layers.IPProtocolTCP):
		c, err = net.Dial("tcp", ipv4Addr)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function net.Dial()", errors.Trace())
		}
	case uint32(layers.IPProtocolUDP):
		c, err = net.Dial("udp", ipv4Addr)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function net.Dial()", errors.Trace())
		}
	}

	return c, nil
}

func portFwdDial(ctx context.Context, payload *mmsp.Payload) error {
	if !auth.AgentMgmtAuth(payload) {
		portFwdDisabled(payload)
		return nil
	}

	var waitc = make(chan struct{})

	proto := payload.PortFwdPDU.PortFwd.Link.Proto
	srcID := payload.PortFwdPDU.PortFwd.Link.SrcNodeID
	dstID := payload.PortFwdPDU.PortFwd.Link.DstNodeID
	dstPort := payload.PortFwdPDU.PortFwd.Link.DstPort
	lID := payload.PortFwdPDU.PortFwd.Link.ID
	cID := payload.PortFwdPDU.PortFwd.Link.ConnectionID

	logging.Debugf("Starting port-forwarding session: srcID %s, dstID %s", srcID, dstID)
	logging.Debugf("Received Dial Request: ConnectionID %s", cID)

	// dial
	conn, err := portFwdConnDial(proto, dstPort)
	if err != nil {
		logging.Errorf("Unable to dial local port %d/%d: %v", proto, dstPort, err)
		portFwdClose(payload)
		return errors.Wrapf(err, "[%v] function portFwdConnListen(ipProto, srcPort)", errors.Trace())
	}

	rpfs.setPortFwdConnection(lID, cID)
	iop := rpfs.getPortFwdConnIO(lID, cID)
	if iop == nil {
		return errors.Errorf("port-fwd io pipes not found for node %s and connection %s", srcID, cID)
	}
	if iop.Out.RP == nil {
		return errors.Errorf("port-fwd output writer pipe not found for node %s and connection %s", srcID, cID)
	}

	// on dial nodes, srcID becomes DstNodeID and dstID becomes SrcNodeID
	go func() {
		portFwdWriteData(dstID, srcID, cID, payload, iop.Out.RP)
		waitc <- struct{}{}
	}()

	if proto == uint32(layers.IPProtocolTCP) {
		conn.(*net.TCPConn).SetKeepAlive(true)
		conn.(*net.TCPConn).SetKeepAlivePeriod(time.Second * 60)
	}

	go func() {
		if _, err := io.Copy(iop.Out.WP, conn); err != nil {
			// logging.Tracef("io.Copy(iop.Out.WP, conn): %v", err)
		}
	}()
	go func() {
		if _, err := io.Copy(conn, iop.In.RP); err != nil {
			// logging.Tracef("io.Copy(conn, iop.In.RP): %v", err)
		}
		waitc <- struct{}{}
	}()

	p := newPortFwdDialAck(payload)
	queuing.TxControlQueue <- p

	<-waitc

	// Shut down the connection.
	logging.Trace("Closing connection...")
	rpfs.deletePortFwdConnection(lID, cID)
	conn.Close()

	portFwdClose(payload)

	return nil
}
