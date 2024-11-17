package app

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/config"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	var grpcServeroptions []grpc.ServerOption

	grpcServeroptions = append(grpcServeroptions, grpc.Creds(a.serviceProvider.TLSCredentials()))

	if a.serviceProvider.antivirusConfig().useAntivirus {
		grpcServeroptions = append(
			grpcServeroptions,
			grpc.UnaryInterceptor(
				a.serviceProvider.antivirusScanner.UnaryServerInterceptor("photo"),
			),
		)
	}

	a.grpcServer = grpc.NewServer(grpcServeroptions...)

	reflection.Register(a.grpcServer)
	userImpl := a.serviceProvider.UserImpl()
	desc.RegisterUserV1Server(a.grpcServer, userImpl)

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
