package interceptors

import (
    "context"
    "strings"

    "github.com/your_project/internal/adapters"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
)

type AuthInterceptor struct {
    keycloakClient *adapters.KeycloakClient
}

func NewAuthInterceptor(kc *adapters.KeycloakClient) *AuthInterceptor {
    return &AuthInterceptor{
        keycloakClient: kc,
    }
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
    return func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
    ) (interface{}, error) {
        err := a.authorize(ctx)
        if err != nil {
            return nil, err
        }
        return handler(ctx, req)
    }
}

func (a *AuthInterceptor) authorize(ctx context.Context) error {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return status.Errorf(codes.Unauthenticated, "metadata is not provided")
    }

    values := md["authorization"]
    if len(values) == 0 {
        return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
    }

    token := strings.TrimPrefix(values[0], "Bearer ")
    if token == "" {
        return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
    }

    _, err := a.keycloakClient.VerifyToken(ctx, token)
    if err != nil {
        return status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
    }

    return nil
}
