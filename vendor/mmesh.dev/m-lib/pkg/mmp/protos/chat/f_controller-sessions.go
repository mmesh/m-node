package chat

import (
	"sync"

	"mmesh.dev/m-api-go/grpc/resources/messaging"
)

type chatSessionID string

type ChatSession struct {
	ThreadID string
	// ThreadTS      string
	ProviderQueue chan *messaging.ChatMessage
	CloseCh       chan struct{}
}

type chatSessionsMap struct {
	session map[chatSessionID]*ChatSession
	sync.RWMutex
}

func newChatSessionsMap() *chatSessionsMap {
	return &chatSessionsMap{
		session: make(map[chatSessionID]*ChatSession),
	}
}

func newChatSession() *ChatSession {
	return &ChatSession{
		ProviderQueue: make(chan *messaging.ChatMessage, 32),
		CloseCh:       make(chan struct{}),
	}
}

func (csm *chatSessionsMap) setChatSession(csID string) {
	csm.Lock()
	defer csm.Unlock()

	if _, ok := csm.session[chatSessionID(csID)]; ok {
		return
	}

	csm.session[chatSessionID(csID)] = newChatSession()
}

func (csm *chatSessionsMap) deleteChatSession(csID string) {
	csm.Lock()
	defer csm.Unlock()

	delete(csm.session, chatSessionID(csID))
}

func (csm *chatSessionsMap) getChatSession(csID string) *ChatSession {
	csm.RLock()
	defer csm.RUnlock()

	cs, ok := csm.session[chatSessionID(csID)]
	if ok {
		return cs
	}

	return nil
}
