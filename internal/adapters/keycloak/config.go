package keycloak

import (
	"os"

	"github.com/pkg/errors"
)

func LoadKeycloakConfig() (KeycloakConfig, error) {
	issuerURL := os.Getenv("APP_KEYCLOAK_ISSUER_URL")
	clientID := os.Getenv("APP_KEYCLOAK_CLIENT_ID")
	clientSecret := os.Getenv("APP_KEYCLOAK_CLIENT_SECRET")

	if len(issuerURL) == 0 {
		return nil, errors.New("keycloak issuer url not found")
	}

	if len(clientID) == 0 {
		return nil, errors.New("keycloak client id not found")
	}

	if len(clientSecret) == 0 {
		return nil, errors.New("keycloak client secret not found")
	}

	return &keycloakConfig{
		IssuerURL: issuerURL,
		ClientID:  clientID,
		ClientSecret: clientSecret,
	}, nil
}