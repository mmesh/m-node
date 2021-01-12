package mmp

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
)

func mgmtAuth(p *mmsp.Payload) bool {
	switch p.PayloadType {
	case PayloadTypeCommandShellExec:
		if viper.GetBool("agent.management.disableExec") {
			return false
		}

	case PayloadTypeTransferInit:
		if viper.GetBool("agent.management.disableTransfer") {
			return false
		}
	case PayloadTypeTransferData:
		if viper.GetBool("agent.management.disableTransfer") {
			return false
		}

	case PayloadTypePortFwdListen:
		if viper.GetBool("agent.management.disablePortForwarding") {
			return false
		}
	case PayloadTypePortFwdDial:
		if viper.GetBool("agent.management.disablePortForwarding") {
			return false
		}
	}

	if !authPSK(p) {
		return false
	}

	if !authSecurityToken(p) {
		return false
	}

	return true
}

func authPSK(p *mmsp.Payload) bool {
	psk := viper.GetString("agent.management.auth.psk")

	if len(psk) > 0 {
		if psk != p.PSK {
			return false
		}
	}

	return true
}

func authSecurityToken(p *mmsp.Payload) bool {
	token := viper.GetString("agent.management.auth.securityToken")

	if len(token) > 0 {
		if token != p.SecurityToken {
			return false
		}
	}

	return true
}
