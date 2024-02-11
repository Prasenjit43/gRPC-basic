package main

import (
	"context"
	"fmt"
	proto "grpc/protoc"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ExampleClient

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client = proto.NewExampleClient(conn)

	r := gin.Default()
	r.GET("/sent", clientConnect)
	r.Run(":8000") //8080
}

func clientConnect(c *gin.Context) {

	req := []*proto.HelloRequest{
		{Something: "one"},
		{Something: "two"},
		{Something: "three"},
		{Something: "four"},
		{Something: "five"},
	}

	stream, err := client.ServerReply(context.TODO())

	for _, x := range req {
		err := stream.Send(x)
		if err != nil {
			return
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
