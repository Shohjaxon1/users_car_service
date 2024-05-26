package service

import (
	"context"
	pb "users_car_service/genproto/user_service"
)

func (s *UserRepo) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	return s.storage.User().Create(req)
}

func (s *UserRepo) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.User, error) {
	return s.storage.User().GetById(req)
}

func (s *UserRepo) GetByPhoneNumber(ctx context.Context, req *pb.GetByPhoneNumberRequest) (*pb.User, error) {
	return s.storage.User().GetByPhoneNumber(req)
}

func (s *UserRepo) GetUserWithCars(ctx context.Context, req *pb.GetByIdRequest) (*pb.AllUsersWithCar, error) {
	return s.storage.User().GetUserWithCars(req)
}

func (s *UserRepo) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.AllUsers, error) {
	return s.storage.User().GetAll(req)
}

func (s *UserRepo) GetAllUserWithCars(ctx context.Context, req *pb.GetAllRequest) (*pb.AllUsersWithCar, error) {
	return s.storage.User().GetAllUserWithCars(req)
}

func (s *UserRepo) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return s.storage.User().Update(req)
}
func (s *UserRepo) DeleteUser(ctx context.Context, req *pb.GetByIdRequest) (*pb.User, error) {
	return s.storage.User().Delete(req)
}
