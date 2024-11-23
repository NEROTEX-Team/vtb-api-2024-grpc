package config

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
)

const (
	useTLS             = "APP_TLS_USE"
	tlsCAFileEnvName   = "APP_TLS_SERVER_CA_FILE_PATH"
	tlsCertFileEnvName = "APP_TLS_SERVER_CERT_FILE_PATH"
	tlsKeyFileEnvName  = "APP_TLS_SERVER_KEY_FILE_PATH"
)

func LoadTLSCredentials() (credentials.TransportCredentials, error) {
	useTLS := os.Getenv(useTLS)
	if useTLS != "true" {
		return nil, errors.New("tls not enabled")
	}
	tlsCertFile := os.Getenv(tlsCertFileEnvName)
	if len(tlsCertFile) == 0 {
		return nil, errors.New("tls cert file not found")
	}

	tlsKeyFile := os.Getenv(tlsKeyFileEnvName)
	if len(tlsKeyFile) == 0 {
		return nil, errors.New("tls key file not found")
	}

	tlsCAFile := os.Getenv(tlsCAFileEnvName)
	if len(tlsCAFile) == 0 {
		return nil, errors.New("tls ca file not found")
	}

	// Load server certificate and private key
	serverCert, err := tls.LoadX509KeyPair(tlsCertFile, tlsKeyFile)
	if err != nil {
		return nil, err
	}

	// Load certificate of the CA who signed client's certificate
	certPool := x509.NewCertPool()
	pemCerts, err := os.ReadFile(tlsCAFile)
	if err != nil {
		return nil, err
	}
	certPool.AppendCertsFromPEM(pemCerts)

	// Create the TLS configuration
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}
