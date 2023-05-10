package portfwd

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/stream"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
)

func NewPortFwd(authKey *auth.AuthKey, srcPort, dstPort uint32, proto, srcID, dstID string, interactive bool) {
	logging.Debugf("New port-forward requested: srcID %s, dstID %s", srcID, dstID)

	ipProto := uint32(ipnet.IPProtocol(proto))

	p := newPortFwdListen(authKey, ipProto, srcPort, dstPort, srcID, dstID, interactive)
	queuing.TxControlQueue <- p
}

func newPortFwdListen(authKey *auth.AuthKey, ipProto, srcPort, dstPort uint32, srcNodeID, dstNodeID string, interactive bool) *mmsp.Payload {
	mmID := viper.GetString("mm.id")
	pfLinkID := portFwdLinkID(ipProto, srcNodeID, srcPort, dstNodeID, dstPort)

	// port-fwd listen session must be started by srcID, so dstID becomes SrcNodeID

	return &mmsp.Payload{
		SrcID: mmID,
		DstID: srcNodeID,
		Type:  mmsp.PDUType_PORTFWD,
		PortFwdPDU: &mmsp.PortFwdPDU{
			Metadata: &mmsp.StreamMetadata{
				RequesterID: mmID,
				Interactive: interactive,
				AuthKey:     authKey,
				// PSK: viper.GetString("management.auth.psk"),
				// SecurityToken: viper.GetString("management.auth.securityToken"),
			},
			Type: mmsp.PortFwdMsgType_PORTFWD_LISTEN,
			PortFwd: &stream.PortFwd{
				Link: &stream.Link{
					ID:        pfLinkID,
					Proto:     ipProto,
					SrcNodeID: srcNodeID,
					SrcPort:   srcPort,
					DstNodeID: dstNodeID,
					DstPort:   dstPort,
				},
				Data: nil,
			},
		},
	}

	// return &mmsp.Payload{
	// 	SrcID:       mmID,
	// 	DstID:       srcNodeID,
	// 	RequesterID: mmID,
	// 	Interactive: interactive,
	// 	AuthKey:     authKey,
	// 	// PSK: viper.GetString("management.auth.psk"),
	// 	// SecurityToken: viper.GetString("management.auth.securityToken"),
	// 	PayloadType: mmsp.PayloadType_PORTFWD_LISTEN,
	// 	PortFwd: &portFwd.PortFwd{
	// 		Link: &portFwd.Link{
	// 			ID:        pfLinkID,
	// 			Proto:     ipProto,
	// 			SrcNodeID: srcNodeID,
	// 			SrcPort:   srcPort,
	// 			DstNodeID: dstNodeID,
	// 			DstPort:   dstPort,
	// 		},
	// 		Data: nil,
	// 	},
	// }
}
