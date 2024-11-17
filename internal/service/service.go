package service

import (
	"context"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/model"
)

type UserService interface {
	CreateUser(ctx context.Context, userData *model.CreateUser) (*model.User, error)
	FetchUserById(ctx context.Context, userID string) (*model.User, error)
	FetchUserList(ctx context.Context, params *model.UserListParams) (*model.UserList, error)
}
