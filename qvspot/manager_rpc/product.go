package manager_rpc

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

// ProductSave creates a product
func (s *managerRPCServer) ProductSave(ctx context.Context, product *qvspot.Product) (*qvspot.Product, error) {

	if _, err := s.qvStore.VendorGetById(ctx, product.VendorId); err != nil {
		if err == store.ErrNotFound {
			return nil, status.Errorf(codes.InvalidArgument, "vendor_id not found")
		}
		return nil, status.Errorf(codes.Internal, "could not fetch vendor", err)
	}

	err := s.qvStore.ProductSave(ctx, product)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not save product: %v", err)
	}

	return product, nil

}

// ProductGet returns the product
func (s *managerRPCServer) ProductGetById(ctx context.Context, request *qvspot.Request) (*qvspot.Product, error) {

	product, err := s.qvStore.ProductGetById(ctx, request.Id)
	if err == store.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "not found")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get product: %v", err)
	}

	return product, nil

}

// ProductDelete deletes a product
func (s *managerRPCServer) ProductDeleteById(ctx context.Context, request *qvspot.Request) (*emptypb.Empty, error) {

	err := s.qvStore.ProductDeleteById(ctx, request.Id)
	if err == store.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "not found")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete product: %v", err)
	}

	return &emptypb.Empty{}, nil

}
