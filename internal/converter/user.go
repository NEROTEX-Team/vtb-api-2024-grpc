package converter

import (
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/model"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreateUserFromDesc(user *desc.CreateUserRequest) *model.CreateUser {
	return &model.CreateUser{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt != nil {
		updatedAt = timestamppb.New(*user.UpdatedAt)
	}

	return &desc.User{
		Id:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
