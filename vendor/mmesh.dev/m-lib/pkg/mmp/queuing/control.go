package queuing

import "mmesh.dev/m-api-go/grpc/network/mmsp"

var RxControlQueue = make(chan *mmsp.Payload, 128)
var TxControlQueue = make(chan *mmsp.Payload, 128)
