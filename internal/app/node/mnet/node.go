package mnet

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmnp/register"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-lib/pkg/errors"
)

func (ln *localNode) NetworkNodeWithoutEndpoints() *network.Node {
	return &network.Node{
		AccountID: ln.accountID,
		TenantID:  ln.tenantID,
		NetID:     ln.netID,
		VRFID:     ln.vrfID,
		NodeID:    ln.nodeID,
	}
}

func (ln *localNode) NetworkNode() *network.Node {
	return &network.Node{
		AccountID:  ln.accountID,
		TenantID:   ln.tenantID,
		NetID:      ln.netID,
		VRFID:      ln.vrfID,
		NodeID:     ln.nodeID,
		Agent:      ln.agent,
		Endpoints:  ln.endpoints.endpt,
		ReplicaSet: ln.replicaSet,
	}
}

func (ln *localNode) registerNode() error {
	nrReq := &register.NodeRegRequest{
		Node: ln.NetworkNode(),
	}

	_, err := ln.Connection().NetworkClient().RegisterNode(context.TODO(), nrReq)
	if err != nil {
		return errors.Wrapf(err, "[%v] function ln.Connection().NetworkClient().RegisterNode()", errors.Trace())
	}

	return nil
}
