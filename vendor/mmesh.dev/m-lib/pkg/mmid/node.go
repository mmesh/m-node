package mmid

import (
	"fmt"
	"strings"

	"mmesh.dev/m-api-go/grpc/network/resources/network"
)

type MMNodeID string

func NewMMNodeID(accountID, tenantID, netID, vrfID, nodeID string) MMID {
	return MMNodeID(fmt.Sprintf("%s:%s:%s:%s:%s", accountID, tenantID, netID, vrfID, nodeID))
}

func GetMMIDFromNode(n *network.Node) MMID {
	return NewMMNodeID(n.AccountID, n.TenantID, n.NetID, n.VRFID, n.NodeID)
}

func (mmid MMNodeID) String() string {
	return string(mmid)
}

func (mmid MMNodeID) AccountID() string {
	return strings.Split(string(mmid), ":")[0]
}

func (mmid MMNodeID) TenantID() string {
	return strings.Split(string(mmid), ":")[1]
}

func (mmid MMNodeID) NetID() string {
	return strings.Split(string(mmid), ":")[2]
}

func (mmid MMNodeID) VRFID() string {
	return strings.Split(string(mmid), ":")[3]
}

func (mmid MMNodeID) VRF() *network.VRF {
	return &network.VRF{
		AccountID: mmid.AccountID(),
		TenantID:  mmid.TenantID(),
		NetID:     mmid.NetID(),
		VRFID:     mmid.VRFID(),
	}
}

func (mmid MMNodeID) NodeID() string {
	return strings.Split(string(mmid), ":")[4]
}

func (mmid MMNodeID) Node() *network.Node {
	return &network.Node{
		AccountID: mmid.AccountID(),
		TenantID:  mmid.TenantID(),
		NetID:     mmid.NetID(),
		VRFID:     mmid.VRFID(),
		NodeID:    mmid.NodeID(),
	}
}
