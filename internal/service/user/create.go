package user

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.UserInfo) (*model.User, error) {
	userID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Failed to generate UUID: %s", err.Error())
		return nil, err
	}

	user, err := s.userRepository.Create(ctx, userID.String(), info)
	if err != nil {
		log.Printf("failed to create user: %s", err.Error())
		return nil, err
	}

	return user, nil
}
