package user

import (
	"context"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/converter"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
)

func (i *Implementation) FetchById(ctx context.Context, req *desc.FetchUserByIdRequest) (*desc.FetchUserByIdResponse, error) {
	user, err := i.userService.FetchUserById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &desc.FetchUserByIdResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
