package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/benny66/protoc/service"
)

func TestChat(t *testing.T) {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatClient(conn)

	user := os.Args[1]
	joinRes, err := client.JoinRoom(context.Background(), &pb.JoinRoomRequest{User: user})
	if err != nil {
		log.Fatalf("could not join room: %v", err)
	}
	log.Println(joinRes.Message)

	for {
		var message string
		fmt.Scanln(&message)

		if message == "exit" {
			break
		}

		_, err := client.SendMessage(context.Background(), &pb.MessageRequest{User: user, Message: message})
		if err != nil {
			log.Printf("could not send message: %v", err)
		}
	}
}
