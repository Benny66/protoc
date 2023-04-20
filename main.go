package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/benny66/protoc/service"
)

type server struct {
	pb.UnimplementedChatServer
}

func (s *server) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("received message from %s: %s", req.User, req.Message)
	return &pb.MessageResponse{Message: fmt.Sprintf("Received message from %s: %s", req.User, req.Message)}, nil
}

func (s *server) JoinRoom(ctx context.Context, req *pb.JoinRoomRequest) (*pb.JoinRoomResponse, error) {
	log.Printf("%s joined the room", req.User)
	return &pb.JoinRoomResponse{Message: fmt.Sprintf("%s joined the room", req.User)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})

	log.Println("starting grpc server on :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
