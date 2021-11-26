package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/resources/messaging"
	"x6a.dev/pkg/xlog"
)

var sessionsMap *chatSessionsMap

var UserMessagesQueue = make(chan *messaging.ChatMessage, 128)

func (a *API) GetChatSession(csID string) *ChatSession {
	if sessionsMap == nil {
		return nil
	}

	return sessionsMap.getChatSession(csID)
}

func newChatExit(srcID string, p *mmsp.Payload) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       p.SrcID,
		PayloadType: mmsp.PayloadType_CHAT_SESSION_EXIT,
	}
}

func (a *API) chatDisabled(payload *mmsp.Payload) {
	mmID := viper.GetString("mm.id")

	p := &mmsp.Payload{
		SrcID:       mmID,
		DstID:       payload.SrcID,
		PayloadType: mmsp.PayloadType_CHAT_SESSION_DISABLED,
	}

	a.TxControlQueue <- p
}

func newChatOperatorOutput(srcID string, p *mmsp.Payload) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       p.SrcID,
		PayloadType: mmsp.PayloadType_CHAT_SESSION_OUTPUT,
		ChatMessage: &messaging.ChatMessage{
			ServiceID:        p.ChatMessage.ServiceID,
			ProviderID:       p.ChatMessage.ProviderID,
			IssueID:          p.ChatMessage.IssueID,
			ThreadID:         p.ChatMessage.ThreadID,
			// ThreadTS:         p.ChatMessage.ThreadTS,
			BackendType:      p.ChatMessage.BackendType,
			Direction:        messaging.ChatMessageDirection_PROVIDER_TO_USER,
			AccountID:        p.ChatMessage.AccountID,
			UserEmail:        p.ChatMessage.UserEmail,
			UserNickname:     p.ChatMessage.UserNickname,
			ProviderChannel:  p.ChatMessage.ProviderChannel,
			OperatorEmail:    p.ChatMessage.OperatorEmail,
			OperatorNickname: p.ChatMessage.OperatorNickname,
			Payload:          nil,
			// Timestamp:
		},
	}
}

func (a *API) chatInit(payload *mmsp.Payload) {
	if sessionsMap == nil {
		sessionsMap = newChatSessionsMap()
	}

	waitc := make(chan struct{})

	mmID := viper.GetString("mm.id")
	// csID := fmt.Sprintf("%s_%s", payload.ChatMessage.ThreadTS, payload.ChatMessage.ThreadID)
	csID := payload.ChatMessage.ThreadID

	sessionsMap.setChatSession(csID)
	cs := sessionsMap.getChatSession(csID)
	if cs == nil {
		xlog.Errorf("chat msg queues not found for session from %s", csID)
		return
	}

	msg := payload.ChatMessage

	go func() {
		a.chatSessionQueueController(mmID, payload, cs)
		waitc <- struct{}{}
	}()

	xlog.Debugf("Chat Session %s (%s) started", csID, msg.UserEmail)

	welcomeMsg := fmt.Sprintf("Chat with %s team started", payload.ChatMessage.ProviderID)

	p := newChatOperatorOutput(mmID, payload)
	p.ChatMessage.OperatorNickname = "m-bot"
	p.ChatMessage.Payload = []byte(welcomeMsg)
	p.ChatMessage.Timestamp = time.Now().Unix()

	a.TxControlQueue <- p

	<-waitc

	// Shut down shell session
	xlog.Debugf("Chat session %s terminated", csID)
	sessionsMap.deleteChatSession(csID)

	// Finish cli session
	p = newChatExit(mmID, payload)
	a.TxControlQueue <- p
}

func chatEnd(payload *mmsp.Payload) {
	// csID := fmt.Sprintf("%s_%s", payload.ChatMessage.ThreadTS, payload.ChatMessage.ThreadID)
	csID := payload.ChatMessage.ThreadID

	if sessionsMap == nil {
		xlog.Errorf("Unable to end chat session %s: not found", csID)
		return
	}

	cs := sessionsMap.getChatSession(csID)
	if cs == nil {
		xlog.Errorf("Unable to end chat missing session %s", csID)
		return
	}

	cs.CloseCh <- struct{}{}
}

func (a *API) chatSessionQueueController(mmID string, payload *mmsp.Payload, cs *ChatSession) {
	for {
		select {
		case msg := <-cs.ProviderQueue:
			p := newChatOperatorOutput(mmID, payload)

			p.ChatMessage.OperatorEmail = msg.OperatorEmail
			p.ChatMessage.OperatorNickname = msg.OperatorNickname
			p.ChatMessage.Payload = msg.Payload
			p.ChatMessage.Timestamp = msg.Timestamp

			a.TxControlQueue <- p
		case <-cs.CloseCh:
			xlog.Debug("Closing chat session queue")
			return
		}
	}
}

func chatReadUserInput(ctx context.Context, payload *mmsp.Payload) error {
	if payload.ChatMessage == nil {
		return nil
	}

	if len(payload.ChatMessage.Payload) == 0 {
		xlog.Trace("len(payload.ChatMessage.Payload) == 0")
		return nil
	}

	UserMessagesQueue <- payload.ChatMessage

	return nil
}
