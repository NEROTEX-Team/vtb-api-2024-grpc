package user

import (
	"context"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/converter"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
)

func (i *Implementation) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*desc.UpdateUserResponse, error) {

	user, err := i.userService.UpdateUserById(ctx, converter.ToUpdateUserFromDesc(req))
	if err != nil {
		return nil, err
	}
	return &desc.UpdateUserResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
