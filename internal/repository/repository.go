package repository

import (
	"context"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, userID string, info *model.UserInfo) error
	Get(ctx context.Context, userID string) (*model.User, error)
}
