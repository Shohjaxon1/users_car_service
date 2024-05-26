package repo

import (
	pb "users_car_service/genproto/user_service"
)

// UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	GetById(*pb.GetByIdRequest) (*pb.User, error)
	GetByPhoneNumber(*pb.GetByPhoneNumberRequest) (*pb.User, error)
	GetAllUserWithCars(*pb.GetAllRequest) (*pb.AllUsersWithCar, error)
	GetUserWithCars(*pb.GetByIdRequest) (*pb.AllUsersWithCar,error)
	GetAll(*pb.GetAllRequest) (*pb.AllUsers, error)
	Update(*pb.User) (*pb.User, error)
	Delete(*pb.GetByIdRequest) (*pb.User, error)
}

type CarStorageI interface {
	CreateCar(*pb.Car) (*pb.Car, error)
	GetCarById(*pb.GetByIdRequest) (*pb.Car, error)
	GetAllCars(*pb.GetAllRequest) (*pb.AllCars, error)
	UpdateCar(*pb.Car) (*pb.Car, error)
	DeleteCar(*pb.GetByIdRequest) (*pb.Car, error)
}