package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	model "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/entities"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
)

func ToCreateUserFromDesc(user *desc.CreateUserRequest) *model.CreateUser {
	return &model.CreateUser{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
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

func ToUsersFromService(users *[]model.User) []*desc.User {
	var result []*desc.User
	for _, user := range *users {
		result = append(result, ToUserFromService(&user))
	}
	return result
}

func ToUpdateUserFromDesc(user *desc.UpdateUserRequest) *model.UpdateUser {
	return &model.UpdateUser{
		ID:        user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
