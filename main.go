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

type Service struct {
	pool map[string]map[string]string
}

func (service *Service) Generate(ctx context.Context, req *pb.UuidRequest) (*pb.UuidResponse, error) {

	// Generate UUID
	uuid, err := Uuid()
	if err != nil {
		return nil, err
	}
	log.Println("Generate UUID: " + uuid)

	// TODO Pool check
	ids, existKey := service.pool[req.Key]
	if !existKey {
		if ids == nil {
			ids = map[string]string{}
		}
		_, exist := ids[uuid]
		if !exist {
			ids[uuid] = ""
			return &pb.UuidResponse{Uuid: uuid}, nil
		}
	}

	// Re Generate UUID
	uuid, err = Uuid()
	if err != nil {
		return nil, err
	}
	log.Println("Re Generate UUID: " + uuid)

	// TODO Re pool check
	ids, existKey = service.pool[req.Key]
	if !existKey {
		if ids == nil {
			ids = map[string]string{}
		}
		_, exist := ids[uuid]
		if !exist {
			ids[uuid] = ""
			return &pb.UuidResponse{Uuid: uuid}, nil
		}
	}
	return nil, err
}

func Uuid() (string, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}
	uuid := string(out[:])
	uuid = strings.TrimRight(uuid, "\n")
	return uuid, nil
}

func main() {
	protcol := "tcp"
	port := ":13009"
	l, err := net.Listen(protcol, port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterUuidGeneratorServer(s, &Service{
		pool: map[string]map[string]string{},
	})
	s.Serve(l)
}
