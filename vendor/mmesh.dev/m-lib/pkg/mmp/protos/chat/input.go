package chat

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
)

// Input runs on controllers
func (a *API) Input(p *mmsp.Payload) error {
	return chatReadUserInput(context.TODO(), p)
}
