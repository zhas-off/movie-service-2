package testutil

import (
	"github.com/zhas-off/movie-service-2/gen"
	"github.com/zhas-off/movie-service-2/rating/internal/controller/rating"
	grpchandler "github.com/zhas-off/movie-service-2/rating/internal/handler/grpc"
	"github.com/zhas-off/movie-service-2/rating/internal/repository/memory"
)

// NewTestRatingGRPCServer creates a new rating gRPC server to be used in tests.
func NewTestRatingGRPCServer() gen.RatingServiceServer {
	r := memory.New()
	ctrl := rating.New(r, nil)
	return grpchandler.New(ctrl)
}
