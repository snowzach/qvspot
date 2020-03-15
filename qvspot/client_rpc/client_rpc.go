package client_rpc

import (
	"context"

	config "github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/snowzach/qvspot/qvspot"
)

type clientRPCServer struct {
	logger  *zap.SugaredLogger
	qvStore qvspot.QVStore

	defaultLimit int32
}

// New returns a new rpc server
func New(qvStore qvspot.QVStore) (qvspot.ClientRPCServer, error) {

	return newServer(qvStore)

}

func newServer(qvStore qvspot.QVStore) (*clientRPCServer, error) {

	return &clientRPCServer{
		logger:  zap.S().With("package", "qvspot.client_rpc"),
		qvStore: qvStore,

		defaultLimit: config.GetInt32("api.default_limit"),
	}, nil

}

// AuthFuncOverride is used if you want to override default authentication for any endpoint
// This disables all authentication for any thingRPC calls
func (s *clientRPCServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}
