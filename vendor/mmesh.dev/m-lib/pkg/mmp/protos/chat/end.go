package chat

import (
	"mmesh.dev/m-api-go/grpc/network/mmsp"
)

// End runs on controllers
func (a *API) End(p *mmsp.Payload) {
	chatEnd(p)
}
