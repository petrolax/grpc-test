package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	hg "github.com/petrolax/grpc-test/hello-grpc"
	u "github.com/petrolax/grpc-test/user"
	"google.golang.org/grpc"
)

type HelloServer struct {
	hg.UnimplementedHelloServiceServer
}

func (bs *HelloServer) SayHello(ctx context.Context, in *hg.HelloRequest) (*hg.HelloReply, error) {
	var user u.User
	_ = json.Unmarshal(in.GetData(), &user)
	log.Println("Received: "+user.Name, user.Surname)
	res, _ := json.Marshal(map[string]string{
		"Result": fmt.Sprintf("Hello, %s %s", user.Name, user.Surname),
	})
	return &hg.HelloReply{Data: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hg.RegisterHelloServiceServer(s, &HelloServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
