package manager

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
)

// VendorInsert creates a vendor
func (s *managerRPCServer) VendorInsert(ctx context.Context, vendor *qvspot.Vendor) (*qvspot.Vendor, error) {

	err := s.qvStore.VendorInsert(vendor)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return vendor, nil

}

// VendorGet returns the vendor
func (s *managerRPCServer) VendorGetById(ctx context.Context, request *qvspot.Request) (*qvspot.Vendor, error) {

	vendor, err := s.qvStore.VendorGetById(request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return vendor, nil

}

// VendorDelete deletes a vendor
func (s *managerRPCServer) VendorDeleteById(ctx context.Context, request *qvspot.Request) (*emptypb.Empty, error) {

	err := s.qvStore.VendorDeleteById(request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return &emptypb.Empty{}, nil

}
