package main

import (
	"log"
	"net"
	"os/exec"
	"strings"

	pb "./uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Service struct{}

func (*Service) Generate(ctx context.Context, req *pb.UuidRequest) (*pb.UuidResponse, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return nil, err
	}
	uuid := string(out[:])
	uuid = strings.TrimRight(uuid, "\n")
	log.Println("Generate UUID: " + uuid)
	return &pb.UuidResponse{Uuid: uuid}, nil
}

func main() {
	protcol := "tcp"
	port := ":13009"
	l, err := net.Listen(protcol, port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterUuidGeneratorServer(s, &Service{})
	s.Serve(l)
}
