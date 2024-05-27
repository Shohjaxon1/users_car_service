package service

import (
	l "student-service/pkg/logger"
	"student-service/storage"

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
