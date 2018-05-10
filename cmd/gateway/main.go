package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/namely/grpc-to-grpc/gen/pb-go/proto"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterJokesHandlerFromEndpoint(ctx, mux, os.Getenv("GRPC_ADDR"), opts)
	if err != nil {
		logrus.WithError(err).Fatal()
	}
	port := os.Getenv("HTTP_PORT")
	logrus.Printf("Starting JSON Gateway server on port %s...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		logrus.WithError(err).Fatal()
	}
}
