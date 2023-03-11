package testutil

import (
	"github.com/zhas-off/movie-service-2/gen"
	"github.com/zhas-off/movie-service-2/movie/internal/controller/movie"
	metadatagateway "github.com/zhas-off/movie-service-2/movie/internal/gateway/metadata/grpc"
	ratinggateway "github.com/zhas-off/movie-service-2/movie/internal/gateway/rating/grpc"
	grpchandler "github.com/zhas-off/movie-service-2/movie/internal/handler/grpc"
	"github.com/zhas-off/movie-service-2/pkg/discovery"
)

// NewTestMovieGRPCServer creates a new movie gRPC server to be used in tests.
func NewTestMovieGRPCServer(registry discovery.Registry) gen.MovieServiceServer {
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	return grpchandler.New(ctrl)
}
