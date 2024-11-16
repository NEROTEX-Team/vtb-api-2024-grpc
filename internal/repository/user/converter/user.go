package converter

import (
	"time"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/model"
	repoModel "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/repository/user/model"
)

func ToUserFromRepo(user *repoModel.User) *model.User {
	var updatedAt *time.Time

	if user.UpdatedAt.Valid {
		updatedAt = &user.UpdatedAt.Time
	}

	return &model.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: updatedAt,
	}
}
