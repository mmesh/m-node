package mmid

import (
	"fmt"
	"strings"

	"mmesh.dev/m-api-go/grpc/resources/network"
)

type MMNetID string

func NewMMNetID(accountID, tenantID, netID string) MMNetID {
	return MMNetID(fmt.Sprintf("%s:%s:%s", accountID, tenantID, netID))
}

func GetMMNetIDFromNode(n *network.Node) MMNetID {
	return MMNetID(fmt.Sprintf("%s:%s:%s", n.AccountID, n.TenantID, n.NetID))
}

func (mmid MMNetID) String() string {
	return string(mmid)
}

func (mmid MMNetID) AccountID() string {
	return strings.Split(string(mmid), ":")[0]
}

func (mmid MMNetID) TenantID() string {
	return strings.Split(string(mmid), ":")[1]
}

func (mmid MMNetID) NetID() string {
	return strings.Split(string(mmid), ":")[2]
}

func (mmid MMNetID) Network() *network.Network {
	return &network.Network{
		AccountID: mmid.AccountID(),
		TenantID:  mmid.TenantID(),
		NetID:     mmid.NetID(),
	}
}
