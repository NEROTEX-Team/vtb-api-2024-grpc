package user

import (
	repositories "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/repositories"
	def "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/services"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repositories.UserRepository
}

func NewService(
	userRepository repositories.UserRepository,
) *service {
	return &service{
		userRepository: userRepository,
	}
}
