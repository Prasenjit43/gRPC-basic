package main

import (
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

func (s *server) ServerReply(req *proto.HelloRequest, stream proto.Example_ServerReplyServer) error {
	fmt.Println("Inside the ServerReply Function")

	resp := []*proto.HelloRespone{
		{Reply: "Hello,"},
		{Reply: "How"},
		{Reply: "are"},
		{Reply: "you?"},
	}

	for _, respMsg := range resp {
		err := stream.Send(respMsg)
		if err != nil {
			return err
		}
		//fmt.Println(respMsg)
	}
	return nil
}
