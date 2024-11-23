package sentry

import (
	"errors"
	"os"
)

type SentryConfig interface {
	UseSentry() bool
	DSN() string
	Environment() string
}

type sentryConfig struct {
	useSentry   bool
	dsn         string
	environment string
}

func (s *sentryConfig) UseSentry() bool {
	return s.useSentry
}

func (s *sentryConfig) DSN() string {
	return s.dsn
}

func (s *sentryConfig) Environment() string {
	return s.environment
}

func LoadSentryConfig() (SentryConfig, error) {
	useSentry := os.Getenv("APP_SENTRY_USE") == "true"
	if !useSentry {
		return &sentryConfig{
			useSentry: false,
		}, nil
	}
	dsn := os.Getenv("APP_SENTRY_DSN")
	if len(dsn) == 0 {
		return nil, errors.New("sentry dsn not found")
	}

	environment := os.Getenv("APP_SENTRY_ENVIRONMENT")
	if len(environment) == 0 {
		environment = "development"
	}

	return &sentryConfig{
		useSentry:   true,
		dsn:         dsn,
		environment: environment,
	}, nil
}
