package storage

import (
	"users_car_service/storage/postgres"
	"users_car_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	User() repo.UserStorageI
	Car()  repo.CarStorageI
}

type storagePg struct {
	db       *sqlx.DB
	userRepo repo.UserStorageI
	carRepo  repo.CarStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
		carRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s storagePg) Car() repo.CarStorageI {
	return s.carRepo
}