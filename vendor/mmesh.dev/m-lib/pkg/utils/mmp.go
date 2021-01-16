package utils

import (
	"fmt"
	"strings"

	"mmesh.dev/m-api-go/grpc/resources/network"
)

func NewMMID(accountID, tenantID, netID, vrfID, nodeID string) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", accountID, tenantID, netID, vrfID, nodeID)
}

func GetAccountFromMMID(mmid string) string {
	return strings.Split(mmid, ":")[0]
}

func GetTenantFromMMID(mmid string) string {
	return strings.Split(mmid, ":")[1]
}

func GetNetworkFromMMID(mmid string) string {
	return strings.Split(mmid, ":")[2]
}

func GetVRFFromMMID(mmid string) string {
	return strings.Split(mmid, ":")[3]
}

func GetNodeIDFromMMID(mmid string) string {
	return strings.Split(mmid, ":")[4]
}

func GetNodeFromMMID(mmid string) *network.Node {
	return &network.Node{
		AccountID: GetAccountFromMMID(mmid),
		TenantID:  GetTenantFromMMID(mmid),
		NetID:     GetNetworkFromMMID(mmid),
		VRFID:     GetVRFFromMMID(mmid),
		NodeID:    GetNodeIDFromMMID(mmid),
	}
}
