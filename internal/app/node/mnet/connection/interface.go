package connection

import (
	"google.golang.org/grpc"
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-lib/pkg/xlog"
)

type Interface interface {
	NetworkClient() rpc.NetworkAPIClient
	Watcher() chan struct{}
	GetExternalIPv4() string
	Close()
}

type session struct {
	connection *connection
	watcherCh  chan struct{}
	endCh      chan struct{}
}

type connection struct {
	grpcClientConn *grpc.ClientConn
	nxnc           rpc.NetworkAPIClient
	externalIPv4   string
}

func New() Interface {
	s := &session{
		connection: newConnection(),
		watcherCh:  make(chan struct{}, 64),
		endCh:      make(chan struct{}, 1),
	}

	go s.connWatcher()

	return s
}

func (s *session) NetworkClient() rpc.NetworkAPIClient {
	if s.connection == nil {
		return nil
	}

	return s.connection.nxnc
}

func (s *session) Watcher() chan struct{} {
	return s.watcherCh
}

func (s *session) GetExternalIPv4() string {
	return s.connection.externalIPv4
}

func (s *session) Close() {
	// ends connection watcher
	s.endCh <- struct{}{}

	// close connection
	if s.connection == nil {
		return
	}

	if err := s.connection.grpcClientConn.Close(); err != nil {
		xlog.Errorf("Unable to close gRPC network connection: %v", err)
	}

	s.connection.nxnc = nil
}
