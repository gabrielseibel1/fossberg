package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/gabrielseibel1/fossberg/protocol"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGameServiceServer
}

func (s *server) EnterGame(ctx context.Context, in *pb.EnterGameRequest) (*pb.EnterGameResponse, error) {
	log.Printf("[EnterGame] %s", in.String())
	return &pb.EnterGameResponse{X: 1, Y: 1, Z: 1}, nil
}

func (s *server) Fire(ctx context.Context, in *pb.FireRequest) (*pb.FireResponse, error) {
	log.Printf("[Fire] %s", in.String())
	return &pb.FireResponse{Hit: true, Dmg: 100}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
