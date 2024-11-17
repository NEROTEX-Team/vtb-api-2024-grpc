package config

import (
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type antivirusConfig struct {
	useAntivirus bool
	url          string
	timeout      time.Duration
	network      string
}

func LoadAntivirusConfig() (*antivirusConfig, error) {
	useAntivirus := os.Getenv("USE_ANTIVIRUS") == "true"
	url := os.Getenv("ANTIVIRUS_URL")
	timeoutStr := os.Getenv("ANTIVIRUS_TIMEOUT")
	if len(timeoutStr) == 0 {
		timeoutStr = "10"
	}
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		return nil, errors.New("antivirus timeout not found")
	}
	network := os.Getenv("ANTIVIRUS_NETWORK")

	if useAntivirus && len(url) == 0 {
		return nil, errors.New("antivirus url not found")
	}

	return &antivirusConfig{
		useAntivirus: useAntivirus,
		url:          url,
		timeout:      time.Duration(timeout) * time.Second,
		network:      network,
	}, nil
}
