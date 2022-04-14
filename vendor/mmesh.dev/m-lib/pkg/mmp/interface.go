package mmp

import (
	"mmesh.dev/m-lib/pkg/mmp/protos/chat"
)

type Interface interface {
}
type API struct {
}

func ChatSession() chat.Interface {
	return &chat.API{
		TxControlQueue: TxControlQueue,
	}
}
