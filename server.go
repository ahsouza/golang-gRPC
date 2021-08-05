package main

import ( 
  "log"
  "net"
  "fmt"
	"google.golang.org/grpc"
  "ahsouza-grpc/chat"
)

const (port = ":8888")

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

  app, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }
  s := chat.Server{}
  
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(app); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}