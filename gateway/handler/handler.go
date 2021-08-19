package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	bg "github.com/petrolax/grpc-test/bye-grpc"
	hg "github.com/petrolax/grpc-test/hello-grpc"
	"google.golang.org/grpc"
)

func Hello(c *gin.Context) {
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}
	defer conn.Close()
	client := hg.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &hg.HelloRequest{Data: bytes})
	if err != nil {
		log.Printf("couldn't not SayHello: %v", err)
	}

	var msg map[string]string
	_ = json.Unmarshal(res.GetData(), &msg)
	c.JSON(http.StatusOK, msg)
}

func Bye(c *gin.Context) {
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}
	defer conn.Close()
	client := bg.NewByeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayBye(ctx, &bg.ByeRequest{Data: bytes})
	if err != nil {
		log.Printf("couldn't not SayBye: %v", err)
	}

	var msg map[string]string
	_ = json.Unmarshal(res.GetData(), &msg)
	c.JSON(http.StatusOK, msg)
}
