//go:generate protoc -I rpc/ rpc/rpc.proto --go_out=plugins=grpc:rpc

package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"

	pb "github.com/robawilkinson/go-file-server/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5051"
)

type server struct{}

func (s *server) GetImage(in *pb.Request, server pb.Files_GetImageServer) error {
	Filename := "./files/" + in.Filename
	data, err := ioutil.ReadFile(Filename)
	buffer := make([]byte, 512)

	file, err := os.Open(Filename)

	if err != nil {
		log.Panic("error", err)
	}
	defer file.Close()

	num, err := file.Read(buffer)
	if err != nil {
		log.Panic("error", err)
	}
	log.Print(num)
	server.Send(&pb.Picture{Image: data, Type: buffer})

	return err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFilesServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
