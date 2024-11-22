package user

import (
	"context"
	"log"

	model "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/entities"
)

func (s *service) UpdateUserById(ctx context.Context, userData *model.UpdateUser) (*model.User, error) {
	user, err := s.userRepository.UpdateUserById(ctx, userData)
	if err != nil {
		log.Printf("failed to update user: %s", err.Error())
		return nil, err
	}

	log.Printf("user was updated: %s", user.ID)

	return user, nil
}
