package netp2p

import (
	"context"
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmnp/register"
	"mmesh.dev/m-api-go/grpc/network/resources/network"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/resources"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

func AddNetworkEndpoint(endpointID, dnsName, reqIPv4 string) (string, error) {
	e := &network.NetworkEndpoint{
		EndpointID: endpointID,
		DNSName:    dnsName,
		ReqIPv4:    reqIPv4,
		// NetworkPolicy: &network.NetworkPolicy{
		// 	DefaultPolicy: getDefaultPolicy(),
		// },
	}

	n := GetNodeWithoutEndpoints()
	erReq := &register.EndpointRegRequest{
		AgentID:  mma.agentID,
		Node:     n,
		Endpoint: e,
		Priority: int32(viper.GetInt("agent.priority")),
	}

	erResp, err := mma.nxnc.RegisterEndpoint(context.Background(), erReq)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function mma.nxnc.RegisterEndpoint()", errors.Trace())
	}

	if erResp.IPv4 == resources.IPAMRequestedIPv4Unavailable {
		xlog.Alertf("Unable to allocate endpoint requested IPv4 %s: subnet is full, no IPv4 address available", e.ReqIPv4)
		os.Exit(1)
	}

	e.IPv4 = erResp.IPv4
	e.IPv6 = erResp.IPv6

	if len(nodeIPv6) == 0 {
		nodeIPv6, err = ipnet.GetIPv6Endpoint(e.IPv6)
		if err != nil {
			xlog.Alertf("Unable to get IPv6 endpoint addr: %s", errors.Cause(err))
			os.Exit(1)
		}
	}

	if err := mma.ifUp(); err != nil {
		xlog.Alertf("Unable to set up interface: %s", errors.Cause(err))
		os.Exit(1)
	}

	if err := ip4AddrAdd(e.IPv4); err != nil {
		xlog.Alertf("Unable to add address %s to interface: %s", e.IPv4, errors.Cause(err))
		os.Exit(1)
	}
	if err := ip6AddrAdd(e.IPv6); err != nil {
		xlog.Alertf("Unable to add address %s to interface: %s", e.IPv4, errors.Cause(err))
		os.Exit(1)
	}

	mma.addNetworkEndpoint(e)

	if err := mma.registerNode(); err != nil {
		return "", errors.Wrapf(err, "[%v] function mma.registerNode()", errors.Trace())
	}

	return e.IPv4, nil
}

func (mma *mmAgent) addNetworkEndpoint(e *network.NetworkEndpoint) {
	mma.endpoints.Lock()
	defer mma.endpoints.Unlock()

	mma.endpoints.endpt[e.EndpointID] = e
}

func RemoveNetworkEndpoint(endpointID string) error {
	e, ok := mma.endpoints.endpt[endpointID]
	if !ok {
		xlog.Warnf("Endpoint %s not found", endpointID)
		return nil
	}

	n := GetNodeWithoutEndpoints()
	erReq := &register.EndpointRegRequest{
		AgentID:  mma.agentID,
		Node:     n,
		Endpoint: e,
		Priority: int32(viper.GetInt("agent.priority")),
	}

	_, err := mma.nxnc.RemoveEndpoint(context.Background(), erReq)
	if err != nil {
		return errors.Wrapf(err, "[%v] function mma.nxnc.RemoveEndpoint()", errors.Trace())
	}

	if err := ip4AddrDel(e.IPv4); err != nil {
		xlog.Alertf("Unable to remove address %s from interface: %s", e.IPv4, errors.Cause(err))
		os.Exit(1)
	}
	if err := ip6AddrDel(e.IPv6); err != nil {
		xlog.Alertf("Unable to remove address %s from interface: %s", e.IPv4, errors.Cause(err))
		os.Exit(1)
	}

	mma.removeNetworkEndpoint(endpointID)

	return nil
}

func (mma *mmAgent) removeNetworkEndpoint(endpointID string) {
	mma.endpoints.Lock()
	defer mma.endpoints.Unlock()

	delete(mma.endpoints.endpt, endpointID)
}
