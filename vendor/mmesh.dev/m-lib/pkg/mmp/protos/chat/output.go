package chat

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
)

// Output runs on user cli
func (a *API) Output(p *mmsp.Payload) error {
	return a.chatReadOperatorOutput(context.TODO(), p)
}
