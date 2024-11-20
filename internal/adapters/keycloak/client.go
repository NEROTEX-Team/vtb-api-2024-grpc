package keycloak

import (
	"context"

	"github.com/coreos/go-oicd"
)

type KeycloakConfig struct {
	IssuerURL string
	ClientID string
}

type KeycloakClient struct {
	verifier *oicd.IDTokenVerifier
}

func NewKeycloakClient(ctx context.Context, config KeycloakConfig) (*KeycloakClient, error) {
	verifier, err := oicd.NewIDTokenVerifier(ctx, config.IssuerURL, config.ClientID)
	if err != nil {
		return nil, err
	}
	return &KeycloakClient{verifier: verifier}, nil
}

func (c *KeycloakClient) VerifyIDToken(ctx context.Context, token string) (*oicd.IDToken, error) {
	return c.verifier.Verify(ctx, token)
}

