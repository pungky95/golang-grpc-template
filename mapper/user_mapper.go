package mapper

import (
	pb "github.com/pungky95/golang-grpc-proto-template/generated"
	"github.com/pungky95/golang-grpc-template/entities"
	"github.com/pungky95/golang-grpc-template/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func EntityToPbUser(user *entities.User) (*pb.User, error) {
	createdAt, err := utils.ConvertToTimestampProto(user.CreatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create user")
	}
	updatedAt, err := utils.ConvertToTimestampProto(user.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create user")
	}

	return &pb.User{
		Id:        user.ID.String(),
		Phone:     user.Phone,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
