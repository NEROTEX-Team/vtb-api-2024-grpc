package app

import (
	"context"
	"log"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/api/user"
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/config"
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/repository"
	userRepository "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/repository/user"
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/service"
	userService "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/service/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/credentials"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	DBCreds string

	pool *pgxpool.Pool

	userRepository repository.UserRepository

	userService service.UserService

	userImpl *user.Implementation

	tlsCredentials credentials.TransportCredentials
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.PGPool())
	}

	return s.userRepository
}

func (s *serviceProvider) UserService() (service.UserService, error) {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository())
	}
	return s.userService, nil
}

func (s *serviceProvider) UserImpl() (*user.Implementation, error) {
	if s.userImpl == nil {
		service, err := s.UserService()
		if err != nil {
			log.Fatalf("failed to get user service: %s", err.Error())
			return nil, err
		}
		s.userImpl = user.NewImplementation(service)
	}

	return s.userImpl, nil
}

func (s *serviceProvider) TLSCredentials() credentials.TransportCredentials {
	if s.tlsCredentials == nil {

		tlsc, err := config.LoadTLSCredentials()
		if err != nil {
			log.Fatalf("failed to get tls credentials: %s", err.Error())
		}

		s.tlsCredentials = tlsc
	}

	return s.tlsCredentials
}

func (s *serviceProvider) DatabaseCredentials() string {
	if s.DBCreds == "" {
		creds, err := config.LoadDatabaseCredentials()
		if err != nil {
			log.Fatalf("failed to get database credentials: %s", err.Error())
		}

		s.DBCreds = creds
	}

	return s.DBCreds
}

func (s *serviceProvider) PGPool() *pgxpool.Pool {
	if s.pool == nil {
		pool, err := pgxpool.New(context.Background(), s.DatabaseCredentials())
		if err != nil {
			log.Fatalf("failed to connect to database: %s", err.Error())
		}
		s.pool = pool
	}

	return s.pool
}
