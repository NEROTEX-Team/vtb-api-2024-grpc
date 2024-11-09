package service

import (
	"context"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfo) (*model.User, error)
	Get(ctx context.Context, uuid string) (*model.User, error)
}
