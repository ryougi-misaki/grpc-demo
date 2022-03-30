package main

import (
	"context"
	"github.com/ryougi-misaki/grpc-demo/middleware"
	protocol "github.com/ryougi-misaki/grpc-demo/protocol"
	"google.golang.org/grpc"
	"log"
	"net"
)

var _ protocol.EchoServiceServer = (*Service)(nil)

type Service struct{
	protocol.UnimplementedEchoServiceServer
}

func (s *Service) Echo(ctx context.Context, req *protocol.EchoRequest) (*protocol.EchoReply, error) {
	res := &protocol.EchoReply{}
	res.Response = req.Request
	return res, nil
}

func main() {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.Auth))

	protocol.RegisterEchoServiceServer(grpcServer,new(Service))

	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(listen)
}