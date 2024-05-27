package service

import (
	"context"
	pb "student-service/genproto/user_service"
)

func (s *UserRepo) CreateCar(ctx context.Context, req *pb.Car) (*pb.Car, error) {
	return s.storage.Car().CreateCar(req)
}

func (s *UserRepo) GetByCarId(ctx context.Context, req *pb.GetByIdRequest) (*pb.Car, error) {
	return s.storage.Car().GetCarById(req)
}

func (s *UserRepo) GetAllCars(ctx context.Context, req *pb.GetAllRequest) (*pb.AllCars, error) {
	return s.storage.Car().GetAllCars(req)
}

func (s *UserRepo) UpdateCar(ctx context.Context, req *pb.Car) (*pb.Car, error) {
	return s.storage.Car().UpdateCar(req)
}
func (s *UserRepo) DeleteCar(ctx context.Context, req *pb.GetByIdRequest) (*pb.Car, error) {
	return s.storage.Car().DeleteCar(req)
}
