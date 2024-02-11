package main

import (
	"context"
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
	r.GET("/", clientConnect)
	r.Run(":8000") //8080

	// req := &proto.HelloRequest{
	// 	Something: "Hi Prasenjit",
	// }

	// client.ServerReply(context.TODO(), req)

}

func clientConnect(c *gin.Context) {
	req := &proto.HelloRequest{
		Something: "Hi Prasenjit",
	}

	client.ServerReply(context.TODO(), req)
	c.JSON(http.StatusOK, gin.H{
		"message": "Run Successfull",
	})
}
