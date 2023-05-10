package client

import (
	"crypto/tls"
	"time"

	"github.com/johnsiilver/getcert"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
)

func newRPCClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (*grpc.ClientConn, error) {
	// Set up the credentials for the connection
	perRPC, err := newRPCCredentials(authKey, authSecret)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function newAPIKey()", errors.Trace())
	}

	grpcDialOpts := []grpc.DialOption{
		// In addition to the following grpc.DialOption, callers may also use
		// the grpc.CallOption grpc.PerRPCCredentials with the RPC invocation
		// itself.
		// See: https://godoc.org/google.golang.org/grpc#PerRPCCredentials
		grpc.WithPerRPCCredentials(perRPC),
		// oauth.NewOauthAccess requires the configuration of transport
		// credentials.
	}

	if viper.GetBool("insecure") {
		grpcDialOpts = append(grpcDialOpts, grpc.WithInsecure())
	} else {
		// tlsCert, xCerts, err := getcert.FromTLSServer(serverEndpoint, true)
		tlsCert, _, err := getcert.FromTLSServer(serverEndpoint, true)
		if err != nil {
			logging.Debug("Connection broken. Invalid TLS Handshake.")
			return nil, errors.Wrapf(err, "[%v] function getcert.FromTLSServer()", errors.Trace())
		}

		// Create tls based credential
		creds := credentials.NewTLS(&tls.Config{
			Certificates:       []tls.Certificate{tlsCert},
			MinVersion:         tls.VersionTLS13,
			NextProtos:         []string{"h2"},
			InsecureSkipVerify: true,
		})

		grpcDialOpts = append(grpcDialOpts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(serverEndpoint, grpcDialOpts...)
	for i := 0; err != nil && i < 30; i++ {
		logging.Trace("Unable to connect to gRPC server, retrying in 3s..")
		time.Sleep(3 * time.Second)
		conn, err = grpc.Dial(serverEndpoint, grpcDialOpts...)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return conn, nil
}

// manager

func NewManagerAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.ManagerAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewManagerAPIClient(conn), conn, nil
}

func NewAccountAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.AccountAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewAccountAPIClient(conn), conn, nil
}

func NewServicesAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.ServicesAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewServicesAPIClient(conn), conn, nil
}

func NewBillingAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.BillingAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewBillingAPIClient(conn), conn, nil
}

// controller

func NewControllerAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.ControllerAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewControllerAPIClient(conn), conn, nil
}

func NewNetworkAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.NetworkAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewNetworkAPIClient(conn), conn, nil
}

func NewIAMAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.IAMAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewIAMAPIClient(conn), conn, nil
}

func NewTenantAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.TenantAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewTenantAPIClient(conn), conn, nil
}

func NewTopologyAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.TopologyAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewTopologyAPIClient(conn), conn, nil
}

func NewMonitoringAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.MonitoringAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewMonitoringAPIClient(conn), conn, nil
}

func NewOpsAPIClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.OpsAPIClient, *grpc.ClientConn, error) {
	conn, err := newRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}

	return rpc.NewOpsAPIClient(conn), conn, nil
}
