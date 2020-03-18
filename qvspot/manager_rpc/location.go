package manager_rpc

import (
	"context"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

// LocationSave creates a location
func (s *managerRPCServer) LocationSave(ctx context.Context, location *qvspot.Location) (*qvspot.Location, error) {

	if _, err := s.qvStore.VendorGetById(ctx, location.VendorId); err != nil {
		if err == store.ErrNotFound {
			return nil, status.Errorf(codes.InvalidArgument, "vendor_id not found")
		}
		return nil, status.Errorf(codes.Internal, "could not fetch vendor", err)
	}

	err := s.qvStore.LocationSave(ctx, location)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not save location: %v", err)
	}

	return location, nil

}

// LocationGet returns the location
func (s *managerRPCServer) LocationGetById(ctx context.Context, request *qvspot.Request) (*qvspot.Location, error) {

	location, err := s.qvStore.LocationGetById(ctx, request.Id)
	if err == store.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "not found")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get location: %v", err)
	}

	return location, nil

}

// LocationDelete deletes a location
func (s *managerRPCServer) LocationDeleteById(ctx context.Context, request *qvspot.Request) (*emptypb.Empty, error) {

	err := s.qvStore.LocationDeleteById(ctx, request.Id)
	if err == store.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "not found")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete location: %v", err)
	}

	return &emptypb.Empty{}, nil

}
