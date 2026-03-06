package main

import (
	_ "cms/srv/user-server/basic/inits"
	"cms/srv/user-server/handler/server"
	"flag"
	"fmt"
	"log"
	"net"

	__ "cms/srv/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

// server is used to implement helloworld.GreeterServer.

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterEcommerceServiceServer(s, &server.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
