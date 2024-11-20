package user

import (
	"context"
	"log"

	model "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/entities"
)

func (s *service) FetchUserById(ctx context.Context, userID string) (*model.User, error) {
	user, err := s.userRepository.FetchUserById(ctx, userID)
	if err != nil {
		log.Printf("failed to get user: %s", err.Error())
		return nil, err
	}
	if user == nil {
		log.Printf("user not found: %s", userID)
		return nil, model.ErrorUserNotFound
	}

	return user, nil
}
