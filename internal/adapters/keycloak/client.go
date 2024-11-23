package keycloak

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

type KeycloakClient struct {
	adminClient    *http.Client
	publicProvider *oidc.Provider
	verifier       *oidc.IDTokenVerifier
	clientID       string
	realm          string
	baseURL        string
}

func NewKeycloakClient(config KeycloakConfig) (*KeycloakClient, error) {
	if !config.UseKeycloak() {
		return nil, nil
	}
	issuerURL := fmt.Sprintf("%s/realms/%s", config.BaseURL(), config.Realm())
	provider, err := oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create OIDC provider: %v", err)
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: config.ClientID(),
	})

	var adminClient *http.Client
	if config.AdminUsername() != "" && config.AdminPassword() != "" {
		adminToken, err := getAdminToken(
			config.BaseURL(),
			config.Realm(),
			config.AdminUsername(),
			config.AdminPassword(),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get admin token: %v", err)
		}

		adminClient = &http.Client{
			Transport: &transportWithToken{
				base:  http.DefaultTransport,
				token: adminToken,
			},
		}
	}

	return &KeycloakClient{
		adminClient:    adminClient,
		publicProvider: provider,
		verifier:       verifier,
		clientID:       config.ClientID(),
		realm:          config.Realm(),
		baseURL:        config.BaseURL(),
	}, nil
}

func getAdminToken(baseURL, realm, username, password string) (string, error) {
	tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", baseURL, realm)

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", "admin-cli")
	data.Set("username", username)
	data.Set("password", password)

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get admin token: %s", string(bodyBytes))
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}

	return tokenResp.AccessToken, nil
}

type transportWithToken struct {
	base  http.RoundTripper
	token string
}

func (t *transportWithToken) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.token)
	return t.base.RoundTrip(req)
}

func (kc *KeycloakClient) VerifyToken(ctx context.Context, token string) (*oidc.IDToken, error) {
	idToken, err := kc.verifier.Verify(ctx, token)
	if err != nil {
		return nil, err
	}
	return idToken, nil
}

func (kc *KeycloakClient) ClientID() string {
	return kc.clientID
}

func (kc *KeycloakClient) CreateUser(ctx context.Context, userData *UserData) error {
	if kc.adminClient == nil {
		return fmt.Errorf("admin client is not configured")
	}

	url := fmt.Sprintf("%s/admin/realms/%s/users", kc.baseURL, kc.realm)

	user := map[string]interface{}{
		"username":  userData.Username,
		"email":     userData.Email,
		"firstName": userData.FirstName,
		"lastName":  userData.LastName,
		"enabled":   true,
		"credentials": []map[string]interface{}{
			{
				"type":      "password",
				"value":     userData.Password,
				"temporary": false,
			},
		},
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := kc.adminClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request to Keycloak: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create user in Keycloak: %s", string(bodyBytes))
	}

	return nil
}

type UserData struct {
	Username  string
	Email     string
	FirstName string
	LastName  string
	Password  string
}
