package cmd

import (
	cli "github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/snowzach/qvspot/conf"
	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/qvspot/product_server"
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

			// Create the GRPC/HTTP server
			s, err := server.New()
			if err != nil {
				logger.Fatalw("Could not create grpc/http server",
					"error", err,
				)
			}

			// Create the rpcserver
			productServer, err := product_server.New(es)
			if err != nil {
				logger.Fatalw("Could not create thing rpcserver",
					"error", err,
				)
			}

			// Register the Thing RPC server to the GRPC Server
			qvspot.RegisterProductRPCServer(s.GRPCServer(), productServer)
			s.GwReg(qvspot.RegisterProductRPCHandlerFromEndpoint)

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
