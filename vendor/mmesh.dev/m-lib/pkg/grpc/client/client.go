package client

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/johnsiilver/getcert"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"mmesh.dev/m-api-go/grpc/network/resources/iam/auth"
	"mmesh.dev/m-api-go/grpc/network/rpc"
	"mmesh.dev/m-lib/pkg/logging"
	"x6a.dev/pkg/errors"
)

var NxClientConn *grpc.ClientConn
var Close = make(chan struct{})

func init() {
	go func() {
		<-Close
		if err := NxClientConn.Close(); err != nil {
			log.Panicln(err)
		}
	}()
}

func newNxRPCClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (*grpc.ClientConn, error) {
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
		//logging.Trace("Unable to connect to gRPC server, retrying in 3s..")
		time.Sleep(3 * time.Second)
		conn, err = grpc.Dial(serverEndpoint, grpcDialOpts...)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}
	//defer conn.Close()

	NxClientConn = conn

	return conn, nil
}

func NewNxNetworkClient(serverEndpoint string, authKey *auth.AuthKey, authSecret string) (rpc.NxNetworkClient, error) {
	conn, err := newNxRPCClient(serverEndpoint, authKey, authSecret)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}
	//defer conn.Close()

	return rpc.NewNxNetworkClient(conn), nil
}

func NewNxAPIClient(serverEndpoint string, authKey *auth.AuthKey) (rpc.NxAPIClient, error) {
	conn, err := newNxRPCClient(serverEndpoint, authKey, "")
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] unable to connect to gRPC server", errors.Trace())
	}
	//defer conn.Close()

	return rpc.NewNxAPIClient(conn), nil
}
