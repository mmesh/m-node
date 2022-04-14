package chat

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmsp"
)

type Interface interface {
	Init(p *mmsp.Payload)
	End(p *mmsp.Payload)
	Exit()
	Input(p *mmsp.Payload) error
	Output(p *mmsp.Payload) error
	Disabled()

	// controller functions
	GetChatSession(csID string) *ChatSession
	chatDisabled(payload *mmsp.Payload)
	chatInit(payload *mmsp.Payload)
	chatSessionQueueController(mmID string, payload *mmsp.Payload, cs *ChatSession)

	// cli functions
	chatReadOperatorOutput(ctx context.Context, p *mmsp.Payload) error
	chatWriteUserInput(ctx context.Context, payload *mmsp.Payload)
	chatEndRequest(srcID string, p *mmsp.Payload)
}
type API struct {
	TxControlQueue chan *mmsp.Payload
}
