package client_rpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
)

// ProductInsert creates a product
func (s *clientRPCServer) ProductSearch(ctx context.Context, search *qvspot.ProductSearchRequest) (*qvspot.ProductSearchResponse, error) {

	if search.Limit == 0 {
		search.Limit = s.defaultLimit
	}

	response, err := s.qvSearch.ProductSearch(ctx, search)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return response, nil

}
