package mnet

import (
	"context"
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmnp/register"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/resources"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (ln *localNode) AddNetworkEndpoint(endpointID, dnsName, reqIPv4 string) (string, error) {
	e := &network.NetworkEndpoint{
		EndpointID: endpointID,
		DNSName:    dnsName,
		ReqIPv4:    reqIPv4,
	}

	erReq := &register.EndpointRegRequest{
		Node:     ln.NetworkNodeWithoutEndpoints(),
		Endpoint: e,
		Priority: int32(viper.GetInt("agent.priority")),
	}

	erResp, err := ln.Connection().NetworkClient().RegisterEndpoint(context.TODO(), erReq)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function ln.Connection().NetworkClient().RegisterEndpoint()", errors.Trace())
	}

	if erResp.IPv4 == resources.IPAMRequestedIPv4Unavailable {
		xlog.Alertf("Unable to allocate endpoint requested IPv4 %s: subnet is full, no IPv4 address available", e.ReqIPv4)
		os.Exit(1)
	}

	e.IPv4 = erResp.IPv4
	e.IPv6 = erResp.IPv6

	if len(ln.Router().IPv4()) == 0 {
		ln.Router().SetIPv4(e.IPv4)
	}

	if len(ln.Router().IPv6()) == 0 {
		ipv6, err := ipnet.GetIPv6Endpoint(e.IPv6)
		if err != nil {
			xlog.Alertf("Unable to get IPv6 endpoint addr: %s", errors.Cause(err))
			os.Exit(1)
		}

		ln.Router().SetIPv6(ipv6)
	}

	if err := ln.Router().IP4AddrAdd(e.IPv4); err != nil {
		xlog.Alertf("Unable to add address %s to interface: %s", e.IPv4, errors.Cause(err))
		os.Exit(1)
	}
	if err := ln.Router().IP6AddrAdd(e.IPv6); err != nil {
		xlog.Alertf("Unable to add address %s to interface: %s", e.IPv4, errors.Cause(err))
		os.Exit(1)
	}

	ln.endpoints.Lock()
	defer ln.endpoints.Unlock()

	ln.endpoints.endpt[e.EndpointID] = e

	if err := ln.registerNode(); err != nil {
		return "", errors.Wrapf(err, "[%v] function localnode.registerNode()", errors.Trace())
	}

	return e.IPv4, nil
}

func (ln *localNode) RemoveNetworkEndpoint(endpointID string) error {
	ln.endpoints.Lock()
	defer ln.endpoints.Unlock()

	e, ok := ln.endpoints.endpt[endpointID]
	if !ok {
		xlog.Warnf("Endpoint %s not found", endpointID)
		return nil
	}

	n := ln.NetworkNodeWithoutEndpoints()
	erReq := &register.EndpointRegRequest{
		Node:     n,
		Endpoint: e,
		Priority: int32(viper.GetInt("agent.priority")),
	}

	_, err := ln.Connection().NetworkClient().RemoveEndpoint(context.TODO(), erReq)
	if err != nil {
		return errors.Wrapf(err, "[%v] function ln.Connection().NetworkClient().RemoveEndpoint()", errors.Trace())
	}

	if err := ln.Router().IP4AddrDel(e.IPv4); err != nil {
		xlog.Alertf("Unable to remove address %s from interface: %s", e.IPv4, errors.Cause(err))
		os.Exit(1)
	}
	if err := ln.Router().IP6AddrDel(e.IPv6); err != nil {
		xlog.Alertf("Unable to remove address %s from interface: %s", e.IPv4, errors.Cause(err))
		os.Exit(1)
	}

	delete(ln.endpoints.endpt, endpointID)

	return nil
}
