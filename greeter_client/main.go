package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	pb "github.com/robawilkinson/go-file-server/rpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

func main() {
	var address string
	if os.Getenv("RPC_ADDRESS") != "" {
		address = os.Getenv("RPC_ADDRESS")
	} else {
		address = "localhost:50051"
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		Filename := strings.Replace(path, "/", "", 1)
		c := pb.NewFilesClient(conn)
		r, err := c.GetImage(context.Background(), &pb.Request{Filename: Filename})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		for {
			response, err := r.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			data := response.Image
			other := response.Type
			w.Header().Set("Content-Type", http.DetectContentType(other))
			w.Write(data)
		}
	})
	log.Fatal(http.ListenAndServe(":3030", nil))
}
