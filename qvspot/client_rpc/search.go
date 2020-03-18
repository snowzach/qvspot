package client_rpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/snowzach/qvspot/qvspot"
)

// Search searches for items or products
func (s *clientRPCServer) Search(ctx context.Context, search *qvspot.SearchRequest) (*qvspot.SearchResponse, error) {

	if search.Limit == 0 {
		search.Limit = s.defaultLimit
	}

	response, err := s.qvSearch.Search(ctx, search)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	return response, nil

}
