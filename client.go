package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Daniel-Fonseca-da-Silva/Chat-GRPC-API/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	connection, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to the server: %v", err)
	}

	defer connection.Close()

	chatClient := chat.NewChatServiceClient(connection)
	stream, err := chatClient.Join(context.Background())
	if err != nil {
		log.Fatalf("Could not join the server: %v", err)
	}

	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	user := scanner.Text()

	// Channel to signal when the client should exit
	done := make(chan bool)

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				if status.Code(err) == codes.Unavailable {
					log.Printf("Server unavailable: %v", err)
				} else {
					log.Printf("Error receiving message: %v", err)
				}
				done <- true
				return
			}

			fmt.Printf("[%s] %s: %s\n", time.Unix(msg.Timestamp, 0).Format("15:04:05"), msg.User, msg.Text)
		}
	}()

	for scanner.Scan() {
		select {
		case <-done:
			fmt.Println("Disconnected from server")
			return
		default:
			msg := &chat.Message{
				User:      user,
				Text:      scanner.Text(),
				Timestamp: time.Now().Unix(),
			}

			if err := stream.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
				return
			}
		}
	}
}
