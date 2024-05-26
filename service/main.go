package service

import (
	l "users_car_service/pkg/logger"
	"users_car_service/storage"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewUserService(db *sqlx.DB, log l.Logger) *UserRepo {
	return &UserRepo{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}
