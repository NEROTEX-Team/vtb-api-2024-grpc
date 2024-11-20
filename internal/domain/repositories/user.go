package repository

import (
	"context"

	model "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, userData *model.CreateUserWithID) (*model.User, error)
	FetchUserById(ctx context.Context, userId string) (*model.User, error)
	FetchUserList(ctx context.Context, params *model.UserListParams) (*[]model.User, error)
	CountUsers(ctx context.Context, params *model.UserListParams) (int64, error)
	UpdateUserById(ctx context.Context, userData *model.UpdateUser) (*model.User, error)
	DeleteUserById(ctx context.Context, userId string) error
	FetchUserByEmail(ctx context.Context, email string) (*model.User, error)
}
