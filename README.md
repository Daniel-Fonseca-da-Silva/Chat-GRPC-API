# Chat gRPC API

A real-time chat application built with gRPC and Go, featuring bidirectional streaming communication between clients and server.

## 🚀 Features

- **Real-time messaging**: Instant message delivery using gRPC bidirectional streaming
- **Multi-client support**: Multiple clients can connect and chat simultaneously
- **Timestamp tracking**: All messages include timestamps for chronological ordering
- **User identification**: Each client can set their username for message attribution
- **Concurrent handling**: Thread-safe server implementation with mutex protection

## 📋 Prerequisites

- Go 1.24.1 or higher
- Protocol Buffers compiler (protoc)
- Go gRPC tools

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Daniel-Fonseca-da-Silva/Chat-GRPC-API.git
   cd Chat-GRPC-API
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Install Protocol Buffers compiler** (if not already installed)
   ```bash
   # Ubuntu/Debian
   sudo apt-get install protobuf-compiler
   
   # macOS
   brew install protobuf
   
   # Windows
   # Download from https://github.com/protocolbuffers/protobuf/releases
   ```

4. **Install Go gRPC tools**
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

## 🔧 Usage

### Starting the Server

1. **Run the server**
   ```bash
   go run server.go
   ```
   The server will start listening on port `50051`.

### Connecting Clients

1. **Run a client**
   ```bash
   go run client.go
   ```

2. **Enter your username** when prompted

3. **Start chatting** - Type your messages and press Enter to send

4. **Run multiple clients** in separate terminals to test multi-user chat

## 📁 Project Structure

```
Chat-GRPC-API/
├── chat/
│   ├── chat.proto          # Protocol Buffers definition
│   ├── chat.pb.go          # Generated Go code for messages
│   └── chat_grpc.pb.go     # Generated Go code for gRPC service
├── server.go               # gRPC server implementation
├── client.go               # gRPC client implementation
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
└── README.md               # This file
```

## 🔌 API Specification

### Protocol Buffers Definition

The chat service is defined in `chat/chat.proto`:

```protobuf
message Message {
    string user = 1;
    string text = 2;
    int64 timestamp = 3;
}

service ChatService {
    rpc Join(stream Message) returns (stream Message);
}
```

### Service Methods

- **Join**: Bidirectional streaming RPC that allows clients to send and receive messages in real-time

## 🏗️ Architecture

- **Server**: Manages client connections and message broadcasting
- **Client**: Connects to server and provides user interface for messaging
- **gRPC**: Handles communication protocol and serialization
- **Protocol Buffers**: Defines message structure and service interface

## 🧪 Testing

To test the application:

1. Start the server in one terminal
2. Start multiple clients in separate terminals
3. Enter different usernames for each client
4. Send messages from any client and observe them appearing in all connected clients

## 🔒 Security Considerations

⚠️ **Note**: This is a basic implementation for demonstration purposes. For production use, consider adding:

- Authentication and authorization
- Message encryption
- Input validation and sanitization
- Rate limiting
- Connection pooling
- Logging and monitoring

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Daniel Fonseca da Silva**

## 🐛 Issues

If you encounter any issues or have suggestions, please [open an issue](https://github.com/Daniel-Fonseca-da-Silva/Chat-GRPC-API/issues) on GitHub. 