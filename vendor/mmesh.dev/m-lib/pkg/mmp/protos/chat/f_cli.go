package chat

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rivo/tview"
	"github.com/spf13/viper"

	// "golang.org/x/crypto/ssh/terminal"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/resources/itsm"
	"mmesh.dev/m-api-go/grpc/resources/messaging"
	"mmesh.dev/m-lib/pkg/mmp/cli"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

var cliStdinWaitc = make(chan struct{}, 1)

var cliStdinActive bool

var UserInputQueue = make(chan string, 64)
var closeUserInputCh = make(chan struct{}, 1)

// func NewChatInit(authKey *auth.AuthKey, issue *itsm.Issue, txControlQueue chan *mmsp.Payload) {
func NewChatInit(issue *itsm.Issue, txControlQueue chan *mmsp.Payload) {
	mmID := viper.GetString("mm.id")

	var threadID string
	var chatBackendType messaging.ChatBackendType
	if issue.Crisp != nil {
		threadID = issue.Crisp.SessionID
		chatBackendType = messaging.ChatBackendType_CRISP
	}

	p := &mmsp.Payload{
		SrcID: mmID,
		// DstID:,
		// AuthKey: authKey,
		PayloadType: mmsp.PayloadType_CHAT_SESSION_INIT,
		ChatMessage: &messaging.ChatMessage{
			ServiceID:  issue.ServiceID,
			ProviderID: issue.ProviderID,
			// ProviderChannel: ,
			IssueID:  issue.IssueID,
			ThreadID: threadID,
			// ThreadTS: ,
			BackendType:  chatBackendType,
			Direction:    messaging.ChatMessageDirection_USER_TO_PROVIDER,
			AccountID:    issue.AccountID,
			UserEmail:    issue.OwnerUserEmail,
			UserNickname: issue.OwnerUserNickname,
			// OperatorEmail: ,
			// OperatorNickname: ,
			Payload:   nil,
			Timestamp: time.Now().Unix(),
		},
	}

	txControlQueue <- p
}

func (a *API) chatEndRequest(srcID string, p *mmsp.Payload) {
	payload := &mmsp.Payload{
		SrcID:       srcID,
		DstID:       p.SrcID,
		PayloadType: mmsp.PayloadType_CHAT_SESSION_END,
		ChatMessage: &messaging.ChatMessage{
			ServiceID:  p.ChatMessage.ServiceID,
			ProviderID: p.ChatMessage.ProviderID,
			IssueID:    p.ChatMessage.IssueID,
			ThreadID:   p.ChatMessage.ThreadID,
			// ThreadTS:         p.ChatMessage.ThreadTS,
			BackendType:      p.ChatMessage.BackendType,
			Direction:        messaging.ChatMessageDirection_USER_TO_PROVIDER,
			AccountID:        p.ChatMessage.AccountID,
			UserEmail:        p.ChatMessage.UserEmail,
			UserNickname:     p.ChatMessage.UserNickname,
			ProviderChannel:  p.ChatMessage.ProviderChannel,
			OperatorEmail:    p.ChatMessage.OperatorEmail,
			OperatorNickname: p.ChatMessage.OperatorNickname,
			Payload:          nil,
			Timestamp:        time.Now().Unix(),
		},
	}

	a.TxControlQueue <- payload
}

func newChatUserInput(srcID string, p *mmsp.Payload) *mmsp.Payload {
	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       p.SrcID,
		PayloadType: mmsp.PayloadType_CHAT_SESSION_INPUT,
		ChatMessage: &messaging.ChatMessage{
			ServiceID:  p.ChatMessage.ServiceID,
			ProviderID: p.ChatMessage.ProviderID,
			IssueID:    p.ChatMessage.IssueID,
			ThreadID:   p.ChatMessage.ThreadID,
			// ThreadTS:         p.ChatMessage.ThreadTS,
			BackendType:      p.ChatMessage.BackendType,
			Direction:        messaging.ChatMessageDirection_USER_TO_PROVIDER,
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

func (a *API) chatReadOperatorOutput(ctx context.Context, p *mmsp.Payload) error {
	if len(p.ChatMessage.Payload) > 0 {
		prefix := tview.TranslateANSI(cli.ChatUserRemote(p.ChatMessage.Timestamp, p.ChatMessage.OperatorNickname))
		msg := fmt.Sprintf("[silver::b]%s", string(p.ChatMessage.Payload))
		fmt.Fprintf(View.conversation, "%s %s\n", prefix, msg)
	}

	go func() {
		if !cliStdinActive {
			a.chatWriteUserInput(ctx, p)
		}
	}()

	return nil
}

func (a *API) chatWriteUserInput(ctx context.Context, payload *mmsp.Payload) {
	cliStdinActive = true

	mmID := viper.GetString("mm.id")

	go func() {
		for {
			select {
			case msg := <-UserInputQueue:
				tm := time.Now().Unix()

				p := newChatUserInput(mmID, payload)
				p.ChatMessage.Payload = []byte(msg)
				p.ChatMessage.Timestamp = tm

				a.TxControlQueue <- p

				prefix := tview.TranslateANSI(cli.ChatUserLocal(tm, payload.ChatMessage.UserNickname))
				msg = fmt.Sprintf("[silver::b]%s", msg)
				fmt.Fprintf(View.conversation, "%s %s\n", prefix, msg)
			case <-closeUserInputCh:
				a.chatEndRequest(mmID, payload)
				fmt.Fprint(View.conversation, "Closing chat session queue\n")
				return
			}
		}
	}()

	<-cliStdinWaitc

	// msg.Debug("Closing standard input..")
}

func chatExit() {
	cliStdinWaitc <- struct{}{}

	time.Sleep(100 * time.Millisecond)

	if View != nil {
		View.Stop()
	}

	// cli.EndOfTransmission()
	cli.Disconnected()

	cliStdinActive = false

	os.Exit(0)
}

func chatUnavailable() {
	text := "Chat Session not available now"

	alert := fmt.Sprintf("%s%s%s", colors.Black("["), colors.Red(text), colors.Black("]"))
	fmt.Println(alert)

	cli.Disconnected()

	os.Exit(0)
}
