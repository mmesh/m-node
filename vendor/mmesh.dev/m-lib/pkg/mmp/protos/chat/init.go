package chat

import (
	"mmesh.dev/m-api-go/grpc/network/mmsp"
)

var chatSessionDisabled bool

// Init runs on controllers
func (a *API) Init(p *mmsp.Payload) {
	if chatSessionDisabled {
		a.chatDisabled(p)
		return
	}

	a.chatInit(p)
}
