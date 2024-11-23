package user

import (
	"context"

	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
)

func (i *Implementation) DeleteUser(ctx context.Context, req *desc.DeleteUserByIdRequest) (*desc.Empty, error) {

	err := i.userService.DeleteUserById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &desc.Empty{}, nil
}
