package netp2p

import (
	"context"

	"mmesh.dev/api-go/grpc/network/mmnp/register"
	"mmesh.dev/api-go/grpc/network/resources/network"
	"x6a.dev/pkg/errors"
)

func GetNodeWithoutEndpoints() *network.Node {
	return &network.Node{
		AccountID: mma.accountID,
		TenantID:  mma.tenantID,
		NetID:     mma.netID,
		VRFID:     mma.vrfID,
		NodeID:    mma.nodeID,
	}
}

func GetNode() *network.Node {
	return mma.getNode()
}

func GetNodeIPv6() string {
	if len(nodeIPv6) == 0 {
		return ""
	}

	return nodeIPv6
}

func (mma *mmAgent) getNode() *network.Node {
	return &network.Node{
		AccountID: mma.accountID,
		TenantID:  mma.tenantID,
		NetID:     mma.netID,
		VRFID:     mma.vrfID,
		NodeID:    mma.nodeID,
		Agent:     mma.agent,
		Endpoints: mma.endpoints.endpt,
	}
}

func (mma *mmAgent) registerNode() error {
	nrReq := &register.NodeRegRequest{
		Node: mma.getNode(),
	}

	_, err := mma.nxnc.RegisterNode(context.Background(), nrReq)
	if err != nil {
		return errors.Wrapf(err, "[%v] function mma.nxnc.RegisterNode()", errors.Trace())
	}

	return nil
}
