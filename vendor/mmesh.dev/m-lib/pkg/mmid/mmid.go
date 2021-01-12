package mmid

import "mmesh.dev/m-api-go/grpc/network/resources/network"

type MMID interface {
	String() string
	AccountID() string
	TenantID() string
	NetID() string
	VRFID() string
	VRF() *network.VRF
}
