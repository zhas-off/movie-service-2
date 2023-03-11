package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/zhas-off/movie-service-2/gen"
	"github.com/zhas-off/movie-service-2/movie/internal/controller/movie"
	metadatagateway "github.com/zhas-off/movie-service-2/movie/internal/gateway/metadata/grpc"
	ratinggateway "github.com/zhas-off/movie-service-2/movie/internal/gateway/rating/grpc"
	grpchandler "github.com/zhas-off/movie-service-2/movie/internal/handler/grpc"
	"github.com/zhas-off/movie-service-2/pkg/discovery"
	"github.com/zhas-off/movie-service-2/pkg/discovery/memory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
)

const serviceName = "movie"

func main() {
	f, err := os.Open("base.yaml")
	if err != nil {
		panic(err)
	}
	var cfg config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	port := cfg.API.Port
	log.Printf("Starting the movie service on port %d", port)
	registry := memory.NewRegistry()
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	defer registry.Deregister(ctx, instanceID, serviceName)
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterMovieServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
