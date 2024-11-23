package interceptors

import (
	"context"
	"strings"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/adapters/keycloak"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	keycloakClient  *keycloak.KeycloakClient
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(kc *keycloak.KeycloakClient) *AuthInterceptor {
	return &AuthInterceptor{
		keycloakClient: kc,
		accessibleRoles: map[string][]string{
			"/user.UserV1/UpdateUser": {"editor"},
			"/user.UserV1/DeleteUser": {"editor"},
			"/user.UserV1/CreateUser": {"editor"},
		},
	}
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		err := a.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (a *AuthInterceptor) authorize(ctx context.Context, method string) error {
	roles, methodProtected := a.accessibleRoles[method]
	if !methodProtected {
		return nil
	}

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

	idToken, err := a.keycloakClient.VerifyToken(ctx, token)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	var claims struct {
		RealmAccess struct {
			Roles []string `json:"roles"`
		} `json:"realm_access"`
		ResourceAccess map[string]struct {
			Roles []string `json:"roles"`
		} `json:"resource_access"`
	}
	if err := idToken.Claims(&claims); err != nil {
		return status.Errorf(codes.Internal, "cannot parse token claims: %v", err)
	}

	userRoles := a.extractUserRoles(claims)

	// Check if user is admin
	if a.isAdmin(userRoles) {
		return nil
	}

	if !a.hasRequiredRole(userRoles, roles) {
		return status.Errorf(codes.PermissionDenied, "permission denied")
	}

	return nil
}

func (a *AuthInterceptor) extractUserRoles(claims struct {
	RealmAccess struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
	ResourceAccess map[string]struct {
		Roles []string `json:"roles"`
	} `json:"resource_access"`
}) []string {
	var userRoles []string
	userRoles = append(userRoles, claims.RealmAccess.Roles...)
	if resource, ok := claims.ResourceAccess[a.keycloakClient.ClientID()]; ok {
		userRoles = append(userRoles, resource.Roles...)
	}
	return userRoles
}

func (a *AuthInterceptor) isAdmin(userRoles []string) bool {
	for _, role := range userRoles {
		if role == "admin" {
			return true
		}
	}
	return false
}

func (a *AuthInterceptor) hasRequiredRole(userRoles, requiredRoles []string) bool {
	roleSet := make(map[string]struct{}, len(userRoles))
	for _, role := range userRoles {
		roleSet[role] = struct{}{}
	}

	for _, requiredRole := range requiredRoles {
		if _, ok := roleSet[requiredRole]; ok {
			return true
		}
	}
	return false
}
