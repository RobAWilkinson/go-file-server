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
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetImage(in *pb.Request, server pb.FileServer_GetImageServer) error {
	// server.Send(&pb.HelloReply{Message: "Hello " + in.Method})
	data, err := ioutil.ReadFile(in.Filename)
	buffer := make([]byte, 512)

	file, err := os.Open(in.Filename)
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
	pb.RegisterFileServerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
