package user

import (
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/adapters/keycloak"
	repositories "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/repositories"
	def "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/services"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repositories.UserRepository
	keycloakClient *keycloak.KeycloakClient
}

func NewService(
	userRepository repositories.UserRepository,
	keycloakClient *keycloak.KeycloakClient,
) *service {
	return &service{
		userRepository: userRepository,
		keycloakClient: keycloakClient,
	}
}
