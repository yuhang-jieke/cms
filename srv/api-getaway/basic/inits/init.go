package inits

import (
	"cms/srv/api-getaway/basic/config"
	__ "cms/srv/proto"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	GrpcInits()
}
func GrpcInits() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	config.ContentClient = __.NewEcommerceServiceClient(conn)
}
