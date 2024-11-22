package user

import (
	"context"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/converter"
	model "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/entities"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
)

func (i *Implementation) FetchUserList(ctx context.Context, req *desc.FetchUserListRequest) (*desc.FetchUserListResponse, error) {
	users, err := i.userService.FetchUserList(ctx, &model.UserListParams{
		Limit:  int64(req.GetLimit()),
		Offset: int64(req.GetOffset()),
	})
	if err != nil {
		return nil, err
	}
	return &desc.FetchUserListResponse{
		Users: converter.ToUsersFromService(&users.Items),
		Total: int32(users.Total),
	}, nil
}
