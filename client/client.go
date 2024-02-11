package main

import (
	"context"
	"fmt"
	proto "grpc/protoc"
	"io"
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

	stream, err := client.ServerReply(context.TODO(), &proto.HelloRequest{
		Something: "Hiiiii",
	})

	if err != nil {
		return
	}

	count := 0

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println("Response Message :", response)
		count++

	}

	c.JSON(http.StatusOK, gin.H{
		"message": count,
	})
}
