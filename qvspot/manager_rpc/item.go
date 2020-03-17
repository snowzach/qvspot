package manager_rpc

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

// ItemInsert creates a item
func (s *managerRPCServer) ItemSave(ctx context.Context, item *qvspot.Item) (*qvspot.Item, error) {

	err := s.qvStore.ItemSave(ctx, item)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	if err = s.PopulateItem(ctx, item); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	if err = s.qvSearch.ItemSave(ctx, item); err != nil {
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

	if err = s.PopulateItem(ctx, item); err != nil {
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

	err = s.qvSearch.ItemDeleteById(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return &emptypb.Empty{}, nil

}

func (s *managerRPCServer) PopulateItem(ctx context.Context, item *qvspot.Item) error {

	var err error

	item.Vendor, err = s.qvStore.VendorGetById(ctx, item.VendorId)
	if err != nil && err != store.ErrNotFound {
		return err
	}

	item.Location, err = s.qvStore.LocationGetById(ctx, item.LocationId)
	if err != nil && err != store.ErrNotFound {
		return err
	}

	item.Product, err = s.qvStore.ProductGetById(ctx, item.ProductId)
	if err != nil && err != store.ErrNotFound {
		return err
	}

	return nil

}
