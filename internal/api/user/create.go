package user

import (
	"context"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/converter"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
)

func (i *Implementation) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {

	user, err := i.userService.CreateUser(ctx, converter.ToCreateUserFromDesc(req))
	if err != nil {
		return nil, err
	}
	return &desc.CreateUserResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
