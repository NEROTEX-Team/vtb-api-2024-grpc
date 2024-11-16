package user

import (
	"context"
	"log"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/model"
)

func (s *service) FetchUserList(ctx context.Context, params *model.UserListParams) (*model.UserList, error) {
	users, err := s.userRepository.FetchUserList(ctx, params)
	if err != nil {
		log.Printf("failed to get user list: %s", err.Error())
		return nil, err
	}

	count, err := s.userRepository.CountUsers(ctx, params)

	return &model.UserList{
		Items: *users,
		Total: count,
	}, nil
}
