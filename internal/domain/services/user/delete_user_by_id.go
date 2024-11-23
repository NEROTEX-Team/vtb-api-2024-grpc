package user

import (
	"context"
	"log"
)

func (s *service) DeleteUserById(ctx context.Context, userID string) error {
	err := s.userRepository.DeleteUserById(ctx, userID)
	if err != nil {
		log.Printf("failed to delete user: %s", err.Error())
		return err
	}

	log.Printf("user was deleted: %s", userID)

	return nil
}
