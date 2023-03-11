package testutil

import (
	"github.com/zhas-off/movie-service-2/gen"
	"github.com/zhas-off/movie-service-2/metadata/internal/controller/metadata"
	grpchandler "github.com/zhas-off/movie-service-2/metadata/internal/handler/grpc"
	"github.com/zhas-off/movie-service-2/metadata/internal/repository/memory"
)

// NewTestMetadataGRPCServer creates a new metadata gRPC server to be used in tests.
func NewTestMetadataGRPCServer() gen.MetadataServiceServer {
	r := memory.New()
	ctrl := metadata.New(r)
	return grpchandler.New(ctrl)
}
