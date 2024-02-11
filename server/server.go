package main

import (
	"context"
	"fmt"
	proto "grpc/protoc"
	"net"


	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")

	if tcpErr != nil {
		panic(tcpErr)
	}

	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) ServerReply(ctx context.Context, req *proto.HelloRequest) (*proto.HelloRespone, error) {
	fmt.Println("Inside the ServerReply Function")
	fmt.Println("Req Data :", req.Something)
	return &proto.HelloRespone{}, nil
}
