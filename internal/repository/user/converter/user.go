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
		Id:        user.Id,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromRepo(info repoModel.UserInfo) model.UserInfo {
	return model.UserInfo{
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Email:     info.Email,
	}
}

func ToUserInfoFromService(info *model.UserInfo) repoModel.UserInfo {
	return repoModel.UserInfo{
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Email:     info.Email,
	}
}
