package mmid

import (
	"fmt"
	"strings"

	"mmesh.dev/m-api-go/grpc/resources/network"
)

type MMVRFID string

func NewMMVRFID(accountID, tenantID, netID, vrfID string) MMID {
	return MMVRFID(fmt.Sprintf("%s:%s:%s:%s", accountID, tenantID, netID, vrfID))
}

func GetMMVRFIDFromNode(n *network.Node) MMVRFID {
	return MMVRFID(fmt.Sprintf("%s:%s:%s:%s", n.AccountID, n.TenantID, n.NetID, n.VRFID))
}

func (mmid MMVRFID) String() string {
	return string(mmid)
}

func (mmid MMVRFID) AccountID() string {
	return strings.Split(string(mmid), ":")[0]
}

func (mmid MMVRFID) TenantID() string {
	return strings.Split(string(mmid), ":")[1]
}

func (mmid MMVRFID) NetID() string {
	return strings.Split(string(mmid), ":")[2]
}

func (mmid MMVRFID) VRFID() string {
	return strings.Split(string(mmid), ":")[3]
}

func (mmid MMVRFID) VRF() *network.VRF {
	return &network.VRF{
		AccountID: mmid.AccountID(),
		TenantID:  mmid.TenantID(),
		NetID:     mmid.NetID(),
		VRFID:     mmid.VRFID(),
	}
}
