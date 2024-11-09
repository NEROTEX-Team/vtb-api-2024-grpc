package converter

import (
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/model"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppd.Timestamp
	if user.UpdatedAt != nil {
		updateAt = timestamppb.New(*user.UpdatedAt)
	}

	return &desc.User{
		Id:        user.UUID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromService(info model.userInfo) *desc.UserInfo {
	return &desc.UserInfo{
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Email:     info.Email,
	}
}

func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Email:     info.Email,
	}
}
