package user

import (
	"sync"

	def "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/repository"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	data map[string]*repoModel.User
	m    sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data: make
	}
}
