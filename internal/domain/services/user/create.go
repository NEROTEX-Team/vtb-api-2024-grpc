package user

import (
	"context"
	"log"

	"github.com/google/uuid"

	model "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/entities"
)

func (s *service) CreateUser(ctx context.Context, userData *model.CreateUser) (*model.User, error) {
	userID, err := uuid.NewUUID()

	kcUserData := &keycloak.UserData{
		Username:  userData.Email,
		Email:     userData.Email,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Password:  userData.Password,
	}

    err := s.keycloakClient.CreateUser(ctx, kcUserData)
    if err != nil {
        return nil, fmt.Errorf("failed to create user in Keycloak: %v", err)
    }

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
