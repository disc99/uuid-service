package main

import (
	"log"
	"net"

	pb "./uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Service struct{}

func (*Service) Generate(ctx context.Context, req *pb.UuidRequest) (*pb.UuidResponse, error) {
	return &pb.UuidResponse{Uuid: "Hello again xxx"}, nil
}

func main() {
	l, err := net.Listen("tcp", ":13009")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterUuidGeneratorServer(s, &Service{})
	s.Serve(l)
}
