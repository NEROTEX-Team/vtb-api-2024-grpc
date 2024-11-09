package converter

import (
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/model"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt != nil {
		updatedAt = timestamppb.New(*user.UpdatedAt)
	}

	return &desc.User{
		ID:        user.ID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromService(info model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Firstname: info.FirstName,
		Lastname:  info.LastName,
		Email:     info.Email,
	}
}

func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		FirstName: info.Firstname,
		LastName:  info.Lastname,
		Email:     info.Email,
	}
}
