package server

import (
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/snowzach/qvspot/embed"
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

	// Serve api-docs and swagger-ui
	fs := http.FileServer(&assetfs.AssetFS{Asset: embed.Asset, AssetDir: embed.AssetDir, AssetInfo: embed.AssetInfo, Prefix: "public"})
	s.router.Get("/api-docs/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		fs.ServeHTTP(w, r)
	}))
	s.router.Get("/swagger-ui/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))

}
