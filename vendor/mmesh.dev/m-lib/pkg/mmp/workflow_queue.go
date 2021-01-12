package mmp

import (
	"mmesh.dev/m-api-go/grpc/network/mmsp"
)

var WorkflowControllerQueue = make(chan *mmsp.Payload, 128)
var WorkflowNodeQueue = make(chan *mmsp.Payload, 128)
