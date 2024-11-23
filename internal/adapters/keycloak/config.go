package keycloak

import (
	"os"

	"github.com/pkg/errors"
)

type KeycloakConfig interface {
	BaseURL() string
	Realm() string
	ClientID() string
	AdminUsername() string
	AdminPassword() string
	UseKeycloak() bool
}

type keycloakConfig struct {
	useKeycloak   bool
	baseURL       string
	realm         string
	clientID      string
	adminUsername string
	adminPassword string
}

func LoadKeycloakConfig() (KeycloakConfig, error) {
	useKeycloak := os.Getenv("APP_KEYCLOAK_ENABLE") == "true"
	if !useKeycloak {
		return &keycloakConfig{
			useKeycloak: false,
		}, nil
	}
	baseURL := os.Getenv("APP_KEYCLOAK_BASE_URL")
	realm := os.Getenv("APP_KEYCLOAK_REALM")
	clientID := os.Getenv("APP_KEYCLOAK_CLIENT_ID")
	adminUsername := os.Getenv("APP_KEYCLOAK_ADMIN_USERNAME")
	adminPassword := os.Getenv("APP_KEYCLOAK_ADMIN_PASSWORD")

	if len(baseURL) == 0 {
		return nil, errors.New("keycloak base url not found")
	}

	if len(realm) == 0 {
		return nil, errors.New("keycloak realm not found")
	}

	if len(clientID) == 0 {
		return nil, errors.New("keycloak client id not found")
	}

	return &keycloakConfig{
		baseURL:       baseURL,
		realm:         realm,
		clientID:      clientID,
		adminUsername: adminUsername,
		adminPassword: adminPassword,
	}, nil
}

func (kc *keycloakConfig) UseKeycloak() bool {
	return kc.useKeycloak
}

func (kc *keycloakConfig) BaseURL() string {
	return kc.baseURL
}

func (kc *keycloakConfig) Realm() string {
	return kc.realm
}

func (kc *keycloakConfig) ClientID() string {
	return kc.clientID
}

func (kc *keycloakConfig) AdminUsername() string {
	return kc.adminUsername
}

func (kc *keycloakConfig) AdminPassword() string {
	return kc.adminPassword
}
