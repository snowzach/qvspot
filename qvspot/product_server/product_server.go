package product_server

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
)

type productRPCServer struct {
	logger       *zap.SugaredLogger
	productStore qvspot.ProductStore
}

// New returns a new rpc server
func New(productStore qvspot.ProductStore) (qvspot.ProductRPCServer, error) {

	return newServer(productStore)

}

func newServer(productStore qvspot.ProductStore) (*productRPCServer, error) {

	return &productRPCServer{
		logger:       zap.S().With("package", "qvspot.product_server"),
		productStore: productStore,
	}, nil

}

// AuthFuncOverride is used if you want to override default authentication for any endpoint
// This disables all authentication for any thingRPC calls
func (s *productRPCServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

// ProductInsert returns all things
func (s *productRPCServer) ProductInsert(ctx context.Context, product *qvspot.Product) (*qvspot.Product, error) {

	s.logger.Infow("Product", zap.Any("product", product))

	err := s.productStore.ProductInsert(product)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return product, nil

}

// ProductDelete deletes a product
func (s *productRPCServer) ProductDelete(ctx context.Context, request *qvspot.ProductId) (*emptypb.Empty, error) {

	err := s.productStore.ProductDeleteById(request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return &emptypb.Empty{}, nil

}
