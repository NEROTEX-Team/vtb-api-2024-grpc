package user

import (
	services "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/services"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService services.UserService
}

func NewImplementation(userService services.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
