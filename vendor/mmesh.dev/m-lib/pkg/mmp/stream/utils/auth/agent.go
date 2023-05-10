package auth

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
)

func AgentMgmtAuth(p *mmsp.Payload) bool {
	var mt *mmsp.StreamMetadata

	switch p.Type {
	case mmsp.PDUType_COMMAND:
		switch p.CommandPDU.Type {
		case mmsp.CommandMsgType_COMMAND_EXEC:
			if viper.GetBool("management.disableExec") {
				return false
			}
		}

		mt = p.CommandPDU.Metadata
	case mmsp.PDUType_TRANSFER:
		switch p.TransferPDU.Type {
		case mmsp.TransferMsgType_TRANSFER_INIT:
			if viper.GetBool("management.disableTransfer") {
				return false
			}
		case mmsp.TransferMsgType_TRANSFER_DATA:
			if viper.GetBool("management.disableTransfer") {
				return false
			}
		}

		mt = p.TransferPDU.Metadata
	case mmsp.PDUType_PORTFWD:
		switch p.PortFwdPDU.Type {
		case mmsp.PortFwdMsgType_PORTFWD_LISTEN:
			if viper.GetBool("management.disablePortForwarding") {
				return false
			}
		case mmsp.PortFwdMsgType_PORTFWD_DIAL:
			if viper.GetBool("management.disablePortForwarding") {
				return false
			}
		}

		mt = p.PortFwdPDU.Metadata
	}

	if mt == nil {
		return true
	}

	if !authPSK(mt) {
		return false
	}

	if !authSecurityToken(mt) {
		return false
	}

	return true
}

func authPSK(mt *mmsp.StreamMetadata) bool {
	psk := viper.GetString("management.auth.psk")

	if len(psk) > 0 {
		if psk != mt.PSK {
			return false
		}
	}

	return true
}

func authSecurityToken(mt *mmsp.StreamMetadata) bool {
	token := viper.GetString("management.auth.securityToken")

	if len(token) > 0 {
		if token != mt.SecurityToken {
			return false
		}
	}

	return true
}
