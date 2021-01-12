package mmp

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/google/gopacket/layers"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/portFwd"
	"mmesh.dev/m-lib/pkg/logging"
	"x6a.dev/pkg/errors"
)

var lpfs = newPortFwdSession()

func newPortFwdDial(cID string, p *mmsp.Payload) *mmsp.Payload {
	// port-fwd dial session must be started by dstID, so dstID becomes DstNodeID
	logging.Debugf("Sending Dial Request: ConnectionID %s", cID)

	return &mmsp.Payload{
		SrcID:         p.PortFwd.Link.SrcNodeID,
		DstID:         p.PortFwd.Link.DstNodeID,
		RequesterID:   p.RequesterID,
		Interactive:   p.Interactive,
		PSK:           viper.GetString("agent.management.auth.psk"),
		SecurityToken: viper.GetString("agent.management.auth.securityToken"),
		PayloadType:   PayloadTypePortFwdDial,
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

func portFwdConnListen(proto, srcPort uint32) (net.Listener, error) {
	var l net.Listener
	var err error
	ipv4Addr := net.JoinHostPort("127.0.0.1", fmt.Sprintf("%d", srcPort))

	switch proto {
	case uint32(layers.IPProtocolTCP):
		l, err = net.Listen("tcp", ipv4Addr)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function net.Listen()", errors.Trace())
		}
	case uint32(layers.IPProtocolUDP):
		l, err = net.Listen("udp", ipv4Addr)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function net.Listen()", errors.Trace())
		}
	}

	return l, nil
}

func portFwdListen(ctx context.Context, payload *mmsp.Payload) error {
	proto := payload.PortFwd.Link.Proto
	srcID := payload.PortFwd.Link.SrcNodeID
	srcPort := payload.PortFwd.Link.SrcPort
	dstID := payload.PortFwd.Link.DstNodeID
	lID := payload.PortFwd.Link.ID

	logging.Debugf("Starting port-forwarding session: srcID %s, dstID %s", srcID, dstID)

	// listen
	l, err := portFwdConnListen(proto, srcPort)
	if err != nil {
		logging.Errorf("Unable open local port %d/%d: %v", proto, srcPort, err)
		return errors.Wrapf(err, "[%v] function portFwdConnListen(ipProto, srcPort)", errors.Trace())
	}

	lpfs.setPortFwdSession(lID)

	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			logging.Error(err)
			if opErr, ok := err.(*net.OpError); ok {
				if !opErr.Temporary() {
					break
				}
			}
			continue
		}

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			waitc := make(chan struct{})

			logging.Trace("New connection...")

			if proto == uint32(layers.IPProtocolTCP) {
				c.(*net.TCPConn).SetKeepAlive(true)
				c.(*net.TCPConn).SetKeepAlivePeriod(time.Second * 60)
			}

			cID := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%v", time.Now().UnixNano())))
			// payload.PortFwd.Link.ConnectionID = cID

			logging.Tracef("New ConnectionID %s for LinkID %s", cID, lID)

			lpfs.setPortFwdConnection(lID, cID)
			iop := lpfs.getPortFwdConnIO(lID, cID)
			if iop == nil {
				logging.Errorf("port-fwd io pipes not found for connection %s", cID)
				return
			}
			if iop.out.rp == nil {
				logging.Errorf("port-fwd output writer pipe not found for connection %s", cID)
				return
			}

			p := newPortFwdDial(cID, payload)
			SendCommandQueue <- p

			logging.Debugf("Waiting for dial ack for connection %s to continue..", cID)
			dialAckCh := lpfs.getPortFwdConnDialAckCh(lID, cID)
			if dialAckCh == nil {
				logging.Errorf("port-fwd dial ack channel not found for connection %s", cID)
				return
			}
			<-*dialAckCh

			// on listen nodes, srcID becomes SrcNodeID and dstID becomes DstNodeID
			go func() {
				portFwdWriteData(srcID, dstID, cID, payload, iop.out.rp)
				waitc <- struct{}{}
			}()

			go func() {
				if _, err := io.Copy(iop.out.wp, c); err != nil {
					//logging.Tracef("io.Copy(iop.out.wp, c): %v", err)
				}
			}()
			go func() {
				if _, err := io.Copy(c, iop.in.rp); err != nil {
					//logging.Tracef("io.Copy(c, iop.in.rp): %v", err)
				}
				waitc <- struct{}{}
			}()

			<-waitc

			// Shut down the connection.
			logging.Trace("Closing connection...")
			lpfs.deletePortFwdConnection(lID, cID)
			conn.Close()
		}(conn)
	}

	if err := l.Close(); err != nil {
		return errors.Wrapf(err, "[%v] function l.Close()", errors.Trace())
	}

	lpfs.deletePortFwdSession(lID)

	// portFwdClose(payload)

	return nil
}

func portFwdDialAck(ctx context.Context, payload *mmsp.Payload) error {
	lID := payload.PortFwd.Link.ID
	cID := payload.PortFwd.Link.ConnectionID

	dialAckCh := lpfs.getPortFwdConnDialAckCh(lID, cID)
	if dialAckCh == nil {
		return errors.Errorf("port-fwd dial ack channel not found for connection %s", cID)
	}
	*dialAckCh <- struct{}{}

	return nil
}
