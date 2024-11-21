package adapters

import (
    "context"

    "github.com/coreos/go-oidc"
)

type KeycloakConfig struct {
    IssuerURL string
    ClientID  string
}

type KeycloakClient struct {
    verifier *oidc.IDTokenVerifier
}

func NewKeycloakClient(cfg KeycloakConfig) *KeycloakClient {
    provider, err := oidc.NewProvider(context.Background(), cfg.IssuerURL)
    if err != nil {
        panic(err)
    }

    verifier := provider.Verifier(&oidc.Config{
        ClientID: cfg.ClientID,
    })

    return &KeycloakClient{
        verifier: verifier,
		ClientID: cfg.ClientID,
    }
}

func (kc *KeycloakClient) VerifyToken(ctx context.Context, token string) (*oidc.IDToken, error) {
    idToken, err := kc.verifier.Verify(ctx, token)
    if err != nil {
        return nil, err
    }
    return idToken, nil
}
