package app

import (
	"context"
	"log"

	antivirus "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/adapters/antivirus"
	database "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/adapters/database"
	userRepository "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/adapters/database/repository/user"
	repository "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/repositories"
	service "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/services"
	userService "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/services/user"
	grpc "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/presentors/grpc"
	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/presentors/grpc/api/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/credentials"
)

type serviceProvider struct {
	grpcConfig     grpc.GRPCConfig
	tlsCredentials credentials.TransportCredentials

	DBCreds string

	pool *pgxpool.Pool

	userRepository repository.UserRepository

	userService service.UserService

	userImpl *user.Implementation

	antivirusConf antivirus.AntivirusConfig

	antivirusScanner *antivirus.Scanner
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() grpc.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := grpc.NewGRPCConfig()
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

		tlsc, err := grpc.LoadTLSCredentials()
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
		creds, err := database.LoadDatabaseCredentials()
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

func (s *serviceProvider) AntivirusConfig() antivirus.AntivirusConfig {
	if s.antivirusConf == nil {
		antivirusConf, err := antivirus.LoadAntivirusConfig()
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
