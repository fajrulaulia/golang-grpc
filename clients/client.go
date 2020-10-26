package main

import (
	"context"
	"gogrpc/blog"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Println("err", err)
	}

	defer conn.Close()

	c := blog.NewBlogServiceClient(conn)
	blog := blog.Blog{
		Title: "berita baru, ini dari client",
		Body:  "HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASBUBYU",
	}

	resp, err := c.SayHello(context.Background(), &blog)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("Response :", resp)

}
