package user

import (
	"context"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/converter"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	user, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}
	return &desc.CreateResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
