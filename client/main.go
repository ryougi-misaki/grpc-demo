package main

import (
	"context"
	"fmt"
	"github.com/ryougi-misaki/grpc-demo/middleware"
	protocol "github.com/ryougi-misaki/grpc-demo/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main()  {
	conn, err := grpc.Dial("localhost:8888",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(middleware.NewAuthentication("admin", "1234567")))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := protocol.NewEchoServiceClient(conn)

	req := &protocol.EchoRequest{Request: "hello"}
	reply, err := client.Echo(context.Background(),req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetResponse())
}