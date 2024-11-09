package user

import (
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/repository"
	def "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

func NewService(
	userRepository repository.UserRepository,
) *service {
	return &service{
		userRepository: userRepository,
	}
}
