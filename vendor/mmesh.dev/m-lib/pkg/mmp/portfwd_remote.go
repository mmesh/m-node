package mmp

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/google/gopacket/layers"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/portFwd"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp/auth"
)

var rpfs = newPortFwdSession()

func newPortFwdDialAck(p *mmsp.Payload) *mmsp.Payload {
	// port-fwd dialAck must be sent by DstNodeID
	logging.Tracef("Sending DialAck: ConnectionID %s", p.PortFwd.Link.ConnectionID)

	return &mmsp.Payload{
		SrcID:       p.PortFwd.Link.DstNodeID,
		DstID:       p.PortFwd.Link.SrcNodeID,
		RequesterID: p.RequesterID,
		Interactive: p.Interactive,
		PayloadType: mmsp.PayloadType_PORTFWD_DIALACK,
		PortFwd: &portFwd.PortFwd{
			Link: &portFwd.Link{
				ID:           p.PortFwd.Link.ID,
				Proto:        p.PortFwd.Link.Proto,
				SrcNodeID:    p.PortFwd.Link.SrcNodeID,
				SrcPort:      p.PortFwd.Link.SrcPort,
				DstNodeID:    p.PortFwd.Link.DstNodeID,
				DstPort:      p.PortFwd.Link.DstPort,
				ConnectionID: p.PortFwd.Link.ConnectionID,
			},
			Data: nil,
		},
	}
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

	proto := payload.PortFwd.Link.Proto
	srcID := payload.PortFwd.Link.SrcNodeID
	dstID := payload.PortFwd.Link.DstNodeID
	dstPort := payload.PortFwd.Link.DstPort
	lID := payload.PortFwd.Link.ID
	cID := payload.PortFwd.Link.ConnectionID

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
			//logging.Tracef("io.Copy(iop.Out.WP, conn): %v", err)
		}
	}()
	go func() {
		if _, err := io.Copy(conn, iop.In.RP); err != nil {
			//logging.Tracef("io.Copy(conn, iop.In.RP): %v", err)
		}
		waitc <- struct{}{}
	}()

	p := newPortFwdDialAck(payload)
	TxControlQueue <- p

	<-waitc

	// Shut down the connection.
	logging.Trace("Closing connection...")
	rpfs.deletePortFwdConnection(lID, cID)
	conn.Close()

	portFwdClose(payload)

	return nil
}
