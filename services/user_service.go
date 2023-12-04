package services

import (
	"github.com/google/uuid"
	pb "github.com/pungky95/golang-grpc-proto-template/generated"
	"github.com/pungky95/golang-grpc-template/config"
	"github.com/pungky95/golang-grpc-template/entities"
	"github.com/pungky95/golang-grpc-template/mapper"
	"github.com/pungky95/golang-grpc-template/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceInterface interface {
	Create(createUserInput *pb.CreateUserInput) (*pb.User, error)
	Update(updateUser *pb.UpdateUserInput) (*pb.User, error)
	Delete(deleteUser *pb.DeleteUserInput) (*pb.User, error)
	Detail(detailUser *pb.DetailUserInput) (*pb.User, error)
}

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserService(
	userRepository repositories.UserRepositoryInterface,
) UserServiceInterface {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s UserService) Create(createUserInput *pb.CreateUserInput) (*pb.User, error) {
	user := &entities.User{
		Phone:    createUserInput.Phone,
		Name:     createUserInput.Name,
		Email:    createUserInput.Email,
		Password: config.GetEnv("DEFAULT_USER_PASSWORD", ""),
	}
	err := s.userRepository.Save(user, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create user")
	}
	return mapper.EntityToPbUser(user)
}

func (s UserService) Update(updateUser *pb.UpdateUserInput) (*pb.User, error) {
	id, err := uuid.Parse(updateUser.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID")
	}
	user, err := s.userRepository.FindOne(entities.User{ID: id}, nil, nil)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	user.Phone = updateUser.Phone
	user.Name = updateUser.Name
	user.Email = updateUser.Email
	err = s.userRepository.Save(&user, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed update user")
	}
	return mapper.EntityToPbUser(&user)
}

func (s UserService) Delete(deleteUser *pb.DeleteUserInput) (*pb.User, error) {
	id, err := uuid.Parse(deleteUser.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID")
	}
	user, err := s.userRepository.FindOne(entities.User{ID: id}, nil, nil)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	err = s.userRepository.Delete(user, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed delete user")
	}
	return mapper.EntityToPbUser(&user)
}

func (s UserService) Detail(detailUser *pb.DetailUserInput) (*pb.User, error) {
	id, err := uuid.Parse(detailUser.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID")
	}
	user, err := s.userRepository.FindOne(entities.User{ID: id}, nil, nil)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return mapper.EntityToPbUser(&user)
}
