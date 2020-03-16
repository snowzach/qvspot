package manager_rpc

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
)

// ItemInsert creates a item
func (s *managerRPCServer) ItemInsert(ctx context.Context, item *qvspot.Item) (*qvspot.Item, error) {

	err := s.qvStore.ItemInsert(ctx, item)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return item, nil

}

// ItemGet returns the item
func (s *managerRPCServer) ItemGetById(ctx context.Context, request *qvspot.Request) (*qvspot.Item, error) {

	item, err := s.qvStore.ItemGetById(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return item, nil

}

// ItemDelete deletes a item
func (s *managerRPCServer) ItemDeleteById(ctx context.Context, request *qvspot.Request) (*emptypb.Empty, error) {

	err := s.qvStore.ItemDeleteById(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return &emptypb.Empty{}, nil

}
