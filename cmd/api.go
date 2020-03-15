package cmd

import (
	cli "github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/snowzach/qvspot/conf"
	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/qvspot/client_rpc"
	"github.com/snowzach/qvspot/qvspot/manager_rpc"
	"github.com/snowzach/qvspot/server"
	"github.com/snowzach/qvspot/store/esearch"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var (
	apiCmd = &cli.Command{
		Use:   "api",
		Short: "Start API",
		Long:  `Start API`,
		Run: func(cmd *cli.Command, args []string) { // Initialize the databse

			// var err error
			// switch config.GetString("storage.type") {
			// case "postgres":
			// 	thingStore, err = postgres.New()
			// }
			// if err != nil {
			// 	logger.Fatalw("Database Error", "error", err)
			// }

			es, err := esearch.New()
			if err != nil {
				logger.Fatalw("Elasticsearch Error", "error", err)
			}
			err = es.Init()
			if err != nil {
				logger.Fatalw("Elasticsearch Init Error", "error", err)
			}

			// Create the GRPC/HTTP server
			s, err := server.New()
			if err != nil {
				logger.Fatalw("Could not create grpc/http server",
					"error", err,
				)
			}

			// Create the rpcserver
			managerRPCServer, err := manager_rpc.New(es)
			if err != nil {
				logger.Fatalw("Could not create manager rpcserver",
					"error", err,
				)
			}

			clientRPCServer, err := client_rpc.New(es)
			if err != nil {
				logger.Fatalw("Could not create client rpcserver",
					"error", err,
				)
			}

			// Register the Thing RPC server to the GRPC Server
			qvspot.RegisterManagerRPCServer(s.GRPCServer(), managerRPCServer)
			s.GwReg(qvspot.RegisterManagerRPCHandlerFromEndpoint)
			qvspot.RegisterClientRPCServer(s.GRPCServer(), clientRPCServer)
			s.GwReg(qvspot.RegisterClientRPCHandlerFromEndpoint)

			err = s.ListenAndServe()
			if err != nil {
				logger.Fatalw("Could not start server",
					"error", err,
				)
			}

			<-conf.Stop.Chan() // Wait until StopChan
			conf.Stop.Wait()   // Wait until everyone cleans up
			zap.L().Sync()     // Flush the logger

		},
	}
)
