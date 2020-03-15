package manager_rpc

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
)

// VendorLocationInsert creates a vendorLocation
func (s *managerRPCServer) VendorLocationInsert(ctx context.Context, vendorLocation *qvspot.VendorLocation) (*qvspot.VendorLocation, error) {

	err := s.qvStore.VendorLocationInsert(ctx, vendorLocation)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return vendorLocation, nil

}

// VendorLocationGet returns the vendorLocation
func (s *managerRPCServer) VendorLocationGetById(ctx context.Context, request *qvspot.Request) (*qvspot.VendorLocation, error) {

	vendorLocation, err := s.qvStore.VendorLocationGetById(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return vendorLocation, nil

}

// VendorLocationDelete deletes a vendorLocation
func (s *managerRPCServer) VendorLocationDeleteById(ctx context.Context, request *qvspot.Request) (*emptypb.Empty, error) {

	err := s.qvStore.VendorLocationDeleteById(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return &emptypb.Empty{}, nil

}
