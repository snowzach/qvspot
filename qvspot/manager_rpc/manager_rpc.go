package manager_rpc

import (
	"context"

	"go.uber.org/zap"

	"github.com/snowzach/qvspot/qvspot"
)

type managerRPCServer struct {
	logger   *zap.SugaredLogger
	qvStore  qvspot.QVStore
	qvSearch qvspot.QVSearch
}

// New returns a new rpc server
func New(qvStore qvspot.QVStore, qvSearch qvspot.QVSearch) (qvspot.ManagerRPCServer, error) {

	return newServer(qvStore, qvSearch)

}

func newServer(qvStore qvspot.QVStore, qvSearch qvspot.QVSearch) (*managerRPCServer, error) {

	return &managerRPCServer{
		logger:   zap.S().With("package", "qvspot.manager_rpc"),
		qvStore:  qvStore,
		qvSearch: qvSearch,
	}, nil

}

// AuthFuncOverride is used if you want to override default authentication for any endpoint
// This disables all authentication for any thingRPC calls
func (s *managerRPCServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}
