package main

import (
	"context"
	"github.com/siqpush/demo-grpc-jg/hello"
	"log"
	"net"
	"google.golang.org/grpc"
)

var count int32 = 0

type myHelloServiceServer struct {
	hello.UnimplementedHelloServiceServer
}

func (receiver myHelloServiceServer) SayHello(context.Context, *hello.HelloRequest) (*hello.HelloReply, error) {
	count++
	return &hello.HelloReply{
		Message: "Hello",
		Count: count,
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myHelloServiceServer{}

	hello.RegisterHelloServiceServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
	
}