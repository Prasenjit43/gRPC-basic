package main

import (
	"fmt"
	proto "grpc/protoc"
	"io"
	"net"
	"strconv"

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

func (s *server) ServerReply(stream proto.Example_ServerReplyServer) error {
	fmt.Println("Inside the ServerReply Function")
	total := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.HelloRespone{
				Reply: strconv.Itoa(total),
			})
		}

		if err != nil {
			return err
		}
		total++
		fmt.Println(req)
	}
}
