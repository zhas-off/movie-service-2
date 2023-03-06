package model

import "github.com/zhas-off/movie-service-2/gen"

// MetadataToProto converts a Metadata struct into a generated proto counterpart.
func MetadataToProto(m *Metadata) *gen.Metadata {
	return &gen.Metadata{
		Id:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Director:    m.Director,
	}
}

// MetadataFromProto converts a generated proto counterpart into a Metadata struct.
func MetadataFromProto(m *gen.Metadata) *Metadata {
	return &Metadata{
		ID:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		Director:    m.Director,
	}
}
