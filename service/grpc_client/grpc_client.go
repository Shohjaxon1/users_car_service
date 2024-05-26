package users_car_serviceClient

import (
	"users_car_service/config"
)

// users_car_serviceClientI ...
type users_car_serviceClientI interface {
}

// users_car_serviceClient ...
type users_car_serviceClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*users_car_serviceClient, error) {
	return &users_car_serviceClient{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}
