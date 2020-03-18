package manager_rpc

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

// VendorSave creates a vendor
func (s *managerRPCServer) VendorSave(ctx context.Context, vendor *qvspot.Vendor) (*qvspot.Vendor, error) {

	err := s.qvStore.VendorSave(ctx, vendor)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not save product: %v", err)
	}

	return vendor, nil

}

// VendorGet returns the vendor
func (s *managerRPCServer) VendorGetById(ctx context.Context, request *qvspot.Request) (*qvspot.Vendor, error) {

	vendor, err := s.qvStore.VendorGetById(ctx, request.Id)
	if err == store.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "not found")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get vendor: %v", err)
	}

	return vendor, nil

}

// VendorDelete deletes a vendor
func (s *managerRPCServer) VendorDeleteById(ctx context.Context, request *qvspot.Request) (*emptypb.Empty, error) {

	err := s.qvStore.VendorDeleteById(ctx, request.Id)
	if err == store.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "not found")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete vendor: %v", err)
	}

	return &emptypb.Empty{}, nil

}
