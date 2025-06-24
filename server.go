package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/Daniel-Fonseca-da-Silva/Chat-GRPC-API/chat"
	"google.golang.org/grpc"
)

type chatServer struct {
	chat.UnimplementedChatServiceServer
	mu       sync.Mutex
	clients  map[chat.ChatService_JoinServer]bool
	messages chan *chat.Message
}

func main() {
	listerner, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	chat.RegisterChatServiceServer(server, newServer())
	server.Serve(listerner)
}

func newServer() *chatServer {
	return &chatServer{
		clients:  make(map[chat.ChatService_JoinServer]bool),
		messages: make(chan *chat.Message),
	}
}

func (s *chatServer) Join(stream chat.ChatService_JoinServer) error {
	s.mu.Lock()
	s.clients[stream] = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, stream)
		s.mu.Unlock()
	}()

	// Goroutine to read messages from the client
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				fmt.Printf("Client disconnected: %v\n", err)
				return
			}
			s.messages <- msg
		}
	}()

	// Loop to send messages to all clients
	for messages := range s.messages {
		s.mu.Lock()
		for client := range s.clients {
			if err := client.Send(messages); err != nil {
				fmt.Printf("Error trying to send messages: %v\n", err)
			}
		}
		s.mu.Unlock()
	}

	return nil
}
