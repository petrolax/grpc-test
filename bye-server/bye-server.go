package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	bg "github.com/petrolax/grpc-test/bye-grpc"
	u "github.com/petrolax/grpc-test/user"
	"google.golang.org/grpc"
)

type ByeServer struct {
	bg.UnimplementedByeServiceServer
}

func (bs *ByeServer) SayBye(ctx context.Context, in *bg.ByeRequest) (*bg.ByeReply, error) {
	var user u.User
	_ = json.Unmarshal(in.GetData(), &user)
	log.Println("Received: "+user.Name, user.Surname)
	res, _ := json.Marshal(map[string]string{
		"Result": fmt.Sprintf("Bye, %s %s", user.Name, user.Surname),
	})
	return &bg.ByeReply{Data: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	bg.RegisterByeServiceServer(s, &ByeServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
