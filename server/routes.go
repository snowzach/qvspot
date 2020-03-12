package server

import (
	"net/http"

	"github.com/snowzach/qvspot/server/versionrpc"
	"github.com/snowzach/qvspot/server/versionrpc/versionrpcserver"
)

// SetupRoutes configures all the routes for this service
func (s *Server) SetupRoutes() {

	// Register our routes - you need at aleast one route
	s.router.Get("/none", func(w http.ResponseWriter, r *http.Request) {})

	// Register RPC Services
	versionrpc.RegisterVersionRPCServer(s.GRPCServer(), versionrpcserver.New())
	s.GwReg(versionrpc.RegisterVersionRPCHandlerFromEndpoint)

}
