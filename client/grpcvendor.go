package main

import (
	"context"
	"crypto/x509"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/snowzach/qvspot/qvspot"
)

func main() {

	// Get the system cert pool
	certPool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalf("Could not get system cert pool: %v", err)
	}

	dialOptions := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(certPool, "")),
	}

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("spot.prozach.org:443", dialOptions...)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	// gRPC version Client
	client := qvspot.NewClientRPCClient(conn)

	request := &qvspot.SearchRequest{
		Search: "product123",
	}

	ctx := context.Background()

	response, err := client.Search(ctx, request)
	log.Printf("Response: %+v err:%v\n", response, err)
}
