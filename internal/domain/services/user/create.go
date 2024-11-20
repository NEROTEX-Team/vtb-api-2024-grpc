package user

import (
	"context"
	"log"

	"github.com/google/uuid"

	model "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/entities"
)

func (s *service) CreateUser(ctx context.Context, userData *model.CreateUser) (*model.User, error) {
	userID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Failed to generate UUID: %s", err.Error())
		return nil, err
	}

	user, err := s.userRepository.CreateUser(ctx, &model.CreateUserWithID{
		ID:        userID.String(),
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
	})
	if err != nil {
		log.Printf("failed to create user: %s", err.Error())
		return nil, err
	}

	return user, nil
}
