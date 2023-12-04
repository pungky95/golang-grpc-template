package server

import (
	"context"

	pb "github.com/pungky95/golang-grpc-proto-template/generated"
	"github.com/pungky95/golang-grpc-template/services"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	userService services.UserServiceInterface
}

func NewUserServer(userService services.UserServiceInterface) *UserServer {
	return &UserServer{userService: userService}
}

func (server *UserServer) CreateUser(_ context.Context, in *pb.CreateUserInput) (*pb.User, error) {
	return server.userService.Create(in)
}

func (server *UserServer) UpdateUser(_ context.Context, in *pb.UpdateUserInput) (*pb.User, error) {
	return server.userService.Update(in)
}

func (server *UserServer) DetailUser(_ context.Context, in *pb.DetailUserInput) (*pb.User, error) {
	return server.userService.Detail(in)
}

func (server *UserServer) DeleteUser(_ context.Context, in *pb.DeleteUserInput) (*pb.User, error) {
	return server.userService.Delete(in)
}
