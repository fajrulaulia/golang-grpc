package main

import (
	"gogrpc/blog"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println("net.Listen", err)
		os.Exit(1)
	}
	grpcServer := grpc.NewServer()

	c := blog.Server{}
	blog.RegisterBlogServiceServer(grpcServer, &c)

	log.Println("GRPC Serve on port 8000")
	if err := grpcServer.Serve(listener); err != nil {
		log.Println("grpcServer.Serve", err)
		os.Exit(1)
	}
}
