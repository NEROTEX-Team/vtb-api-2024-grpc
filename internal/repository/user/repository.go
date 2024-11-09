package user

import (
	"context"
	"sync"
	"time"

	def "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/repository"
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/repository/user/converter"
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/repository/user/model"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	data map[string]*repoModel.User
	m    sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]*repoModel.User),
	}
}

func (r *repository) Create(_ context.Context, userUUID string, info *model.UserInfo) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.data[userUUID] = &repoModel.User{
		UUID:      userUUID,
		Info:      converter.ToUserInfoFromService(info),
		CreatedAt: time.Now(),
	}

	return nil
}

func (r *repository) Get(_ context.Context, uuid string) (*model.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	user, ok := r.data[uuid]
	if !ok {
		return nil, nil
	}

	return converter.ToUserFromRepo(user), nil
}
