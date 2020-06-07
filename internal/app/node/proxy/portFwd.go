package proxy

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/google/gopacket/layers"
	"mmesh.dev/mmesh/internal/pkg/ipnet"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

func connListen(af ipnet.AddressFamily, addr string, proto layers.IPProtocol, srcPort uint32) (net.Listener, error) {
	var l net.Listener
	var err error

	ipAddr := net.JoinHostPort(addr, fmt.Sprintf("%d", srcPort))

	netProto, err := ipnet.GetNetworkProtocol(af, proto)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function ipnet.GetNetworkProtocol()", errors.Trace())
	}

	switch proto {
	case layers.IPProtocolTCP:
		l, err = net.Listen(netProto, ipAddr)
		if err != nil {
			// if opErr, ok := err.(*net.OpError); ok {
			// 	fmt.Println("net.Listen:", opErr.Error())
			// 	litter.Dump(opErr)
			// }

			return nil, errors.Wrapf(err, "[%v] function net.Listen()", errors.Trace())
		}
	case layers.IPProtocolUDP:
		return nil, fmt.Errorf("UDP not implemented")
		// l, err = net.Listen("udp", ipAddr)
		// if err != nil {
		// 	return nil, errors.Wrapf(err, "[%v] function net.Listen()", errors.Trace())
		// }
	}

	return l, nil
}

func portFwd(af ipnet.AddressFamily, svcName, vip, ip, portName string, proto layers.IPProtocol, port int32, quitCh chan struct{}, wg *sync.WaitGroup) error {
	xlog.Debugf("Starting mmesh forwarding for svc %s/%s: %s -> %s (%s/%d)", svcName, portName, vip, ip, proto.String(), port)

	// listen
	l, err := connListen(af, vip, proto, uint32(port))
	if err != nil {
		xlog.Errorf("Unable open local port %v/%d: %v", proto, port, errors.Cause(err))

		wg.Done()

		if opErr, ok := errors.Cause(err).(*net.OpError); ok {
			if !opErr.Temporary() {
				return errors.Wrapf(err, "[%v] function connListen()", errors.Trace())
			}
		}
		return nil
	}

	go func() {
		for {
			// Wait for a connection.
			srcConn, err := l.Accept()
			if err != nil {
				// xlog.Errorf("Unable to accept connection: %v", err)
				if opErr, ok := err.(*net.OpError); ok {
					if !opErr.Temporary() {
						break
					}
				}
				continue
			}

			go func() {
				waitc := make(chan struct{}, 2)

				hostPort := net.JoinHostPort(ip, fmt.Sprintf("%d", port))

				netProto, err := ipnet.GetNetworkProtocol(ipnet.AddressFamilyIPv4, proto)
				if err != nil {
					srcConn.Close()
					xlog.Error(err)
					return
				}

				dstConn, err := net.Dial(netProto, hostPort)
				if err != nil {
					srcConn.Close()
					xlog.Errorf("Unable to dial to service %s: %v", svcName, err)
					return
				}

				if proto == layers.IPProtocolTCP {
					srcConn.(*net.TCPConn).SetKeepAlive(true)
					srcConn.(*net.TCPConn).SetKeepAlivePeriod(time.Second * 60)
				}

				go func() {
					io.Copy(dstConn, srcConn)
					waitc <- struct{}{}
				}()
				go func() {
					io.Copy(srcConn, dstConn)
					waitc <- struct{}{}
				}()

				<-waitc

				srcConn.Close()
				dstConn.Close()
			}()
		}

		quitCh <- struct{}{}
	}()

	<-quitCh

	xlog.Debugf("Closing network connection to %s/%s", svcName, portName)

	if err := l.Close(); err != nil {
		xlog.Errorf("Unable to close listener for service %s: %v", svcName, err)
	}

	wg.Done()

	return nil
}
