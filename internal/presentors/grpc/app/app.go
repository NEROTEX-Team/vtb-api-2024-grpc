package app

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	interceptors "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/presentors/grpc/interceptors"
	desc "github.com/NEROTEX-Team/vtb-api-2024-grpc/pkg/v1/user"

	"github.com/getsentry/sentry-go"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_sentry "github.com/johnbellone/grpc-middleware-sentry"
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

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	var grpcServeroptions []grpc.ServerOption
	if a.serviceProvider.TLSCredentials() == nil {
		grpcServeroptions = append(grpcServeroptions, grpc.Creds(insecure.NewCredentials()))
	} else {
		grpcServeroptions = append(grpcServeroptions, grpc.Creds(a.serviceProvider.TLSCredentials()))
	}

	if a.serviceProvider.AntivirusConfig().UseAntivirus() {
		grpcServeroptions = append(
			grpcServeroptions,
			grpc.UnaryInterceptor(
				interceptors.AntivirusInterceptor(a.serviceProvider.antivirusScanner, "photo"),
			),
		)
	}

	if a.serviceProvider.KeycloakConfig().UseKeycloak() {
		grpcServeroptions = append(
			grpcServeroptions,
			grpc.UnaryInterceptor(
				interceptors.NewAuthInterceptor(a.serviceProvider.KeycloakClient()).Unary(),
			),
		)
	}

	if a.serviceProvider.SentryConfig().UseSentry() {
		log.Print("Sentry enabled")
		err := sentry.Init(sentry.ClientOptions{
			Dsn:          a.serviceProvider.SentryConfig().DSN(),
			Debug:        a.serviceProvider.SentryConfig().Environment() != "production",
			Environment:  a.serviceProvider.SentryConfig().Environment(),
			Release:      "vtb-api-hack-server@1.0.0",
			IgnoreErrors: []string{},
		})
		defer sentry.Flush(2 * time.Second)
		if err != nil {
			log.Fatal(err.Error())
		}
		grpcServeroptions = append(
			grpcServeroptions,
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_sentry.UnaryServerInterceptor(),
			)),
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
