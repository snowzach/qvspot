package manager_rpc

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
)

// ProductLocationInsert creates a productLocation
func (s *managerRPCServer) ProductLocationInsert(ctx context.Context, productLocation *qvspot.ProductLocation) (*qvspot.ProductLocation, error) {

	err := s.qvStore.ProductLocationInsert(ctx, productLocation)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return productLocation, nil

}

// ProductLocationGet returns the productLocation
func (s *managerRPCServer) ProductLocationGetById(ctx context.Context, request *qvspot.Request) (*qvspot.ProductLocation, error) {

	productLocation, err := s.qvStore.ProductLocationGetById(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return productLocation, nil

}

// ProductLocationDelete deletes a productLocation
func (s *managerRPCServer) ProductLocationDeleteById(ctx context.Context, request *qvspot.Request) (*emptypb.Empty, error) {

	err := s.qvStore.ProductLocationDeleteById(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return &emptypb.Empty{}, nil

}
