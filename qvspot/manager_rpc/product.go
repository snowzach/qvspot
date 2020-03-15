package manager_rpc

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
)

// ProductInsert creates a product
func (s *managerRPCServer) ProductInsert(ctx context.Context, product *qvspot.Product) (*qvspot.Product, error) {

	err := s.qvStore.ProductInsert(ctx, product)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return product, nil

}

// ProductGet returns the product
func (s *managerRPCServer) ProductGetById(ctx context.Context, request *qvspot.Request) (*qvspot.Product, error) {

	product, err := s.qvStore.ProductGetById(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return product, nil

}

// ProductDelete deletes a product
func (s *managerRPCServer) ProductDeleteById(ctx context.Context, request *qvspot.Request) (*emptypb.Empty, error) {

	err := s.qvStore.ProductDeleteById(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return &emptypb.Empty{}, nil

}
