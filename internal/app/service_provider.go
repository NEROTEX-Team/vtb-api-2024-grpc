package app

import (
	"context"
	"log"

	antivirus "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/adapters/antivirus"
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
	grpcConfig     config.GRPCConfig
	tlsCredentials credentials.TransportCredentials

	DBCreds string

	pool *pgxpool.Pool

	userRepository repository.UserRepository

	userService service.UserService

	userImpl *user.Implementation

	antivirusConf config.AntivirusConfig

	antivirusScanner *antivirus.Scanner
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

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository())
	}
	return s.userService
}

func (s *serviceProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService())
		log.Print("User impl created")
	}

	return s.userImpl
}

func (s *serviceProvider) TLSCredentials() credentials.TransportCredentials {
	if s.tlsCredentials == nil {

		tlsc, err := config.LoadTLSCredentials()
		if err != nil {
			log.Printf("failed to get tls credentials: %s", err.Error())
			return nil
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

func (s *serviceProvider) AntivirusConfig() config.AntivirusConfig {
	if s.antivirusConf == nil {
		antivirusConf, err := config.LoadAntivirusConfig()
		if err != nil {
			log.Fatalf("failed to load antivirus config: %s", err.Error())
		}
		s.antivirusConf = antivirusConf
	}
	return s.antivirusConf
}

func (s *serviceProvider) AntivirusScanner() *antivirus.Scanner {
	if s.antivirusScanner == nil {
		conf := s.AntivirusConfig()
		s.antivirusScanner = antivirus.NewScanner(
			conf.Address(),
			conf.Network(),
			conf.Timeout(),
			conf.UseAntivirus(),
		)
	}
	return s.antivirusScanner
}
