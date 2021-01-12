package client

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/grpc/credentials"
	"mmesh.dev/m-api-go/grpc/network/resources/iam/auth"
	"x6a.dev/pkg/errors"
)

type RPCCredentials struct {
	auth.AuthKey
	AuthSecret string
}

// newRPCCredentials gets the authorization bearer string key
func newRPCCredentials(authKey *auth.AuthKey, authSecret string) (credentials.PerRPCCredentials, error) {
	if len(authKey.Key) > 0 {
		return &RPCCredentials{
			AuthKey:    *authKey,
			AuthSecret: authSecret,
		}, nil
	}

	return nil, errors.Errorf("invalid or inexistent authKey")
}

func (c *RPCCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"key":    c.Key,
		"secret": c.AuthSecret,
	}, nil
}

func (c *RPCCredentials) RequireTransportSecurity() bool {
	if viper.GetBool("insecure") {
		return false
	}

	return true
}
