package grpc_client

import (
	"student-service/config"
)

type StudentServiceClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*StudentServiceClient, error) {
	return &StudentServiceClient{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}
