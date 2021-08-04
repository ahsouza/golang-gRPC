package main

import ( 
  "log"
  "net"
	"google.golang.org/grpc"
)

func main() {
  app, err := net.Listen("tcp", ":8888")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(app); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}