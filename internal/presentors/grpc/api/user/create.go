package user

import (
	"context"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/converter"
	model "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/entities"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {

	user, err := i.userService.CreateUser(ctx, converter.ToCreateUserFromDesc(req))
	if err != nil {
		if err == model.ErrorUserAlreadyExists {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, err
	}
	return &desc.CreateUserResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
