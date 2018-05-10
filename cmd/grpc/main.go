package main

import (
	"context"
	"fmt"
	"net"
	"os"

	pb "github.com/namely/grpc-to-grpc/gen/pb-go/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetJoke(ctx context.Context, req *pb.JokeReq) (*pb.JokeResp, error) {
	return &pb.JokeResp{
		Body: "This is a joke right?",
	}, nil
}

func main() {
	port := os.Getenv("GRPC_PORT")

	logrus.Printf("Starting RPC server on port %s...", port)
	list, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	s := grpc.NewServer()
	pb.RegisterJokesServer(s, &server{})

	if err := s.Serve(list); err != nil {
		logrus.WithError(err).Fatal()
	}
}
