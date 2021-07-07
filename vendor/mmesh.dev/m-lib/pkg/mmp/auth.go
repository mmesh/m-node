package mmp

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
)

func mgmtAuth(p *mmsp.Payload) bool {
	switch p.PayloadType {
	case mmsp.PayloadType_COMMAND_SHELL_EXEC:
		if viper.GetBool("agent.management.disableExec") {
			return false
		}

	case mmsp.PayloadType_TRANSFER_INIT:
		if viper.GetBool("agent.management.disableTransfer") {
			return false
		}
	case mmsp.PayloadType_TRANSFER_DATA:
		if viper.GetBool("agent.management.disableTransfer") {
			return false
		}

	case mmsp.PayloadType_PORTFWD_LISTEN:
		if viper.GetBool("agent.management.disablePortForwarding") {
			return false
		}
	case mmsp.PayloadType_PORTFWD_DIAL:
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
