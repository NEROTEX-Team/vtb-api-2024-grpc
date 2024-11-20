package antivirus

import (
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type AntivirusConfig interface {
	UseAntivirus() bool
	Address() string
	Timeout() time.Duration
	Network() string
}

type antivirusConfig struct {
	useAntivirus bool
	address      string
	timeout      time.Duration
	network      string
}

func LoadAntivirusConfig() (AntivirusConfig, error) {
	useAntivirus := os.Getenv("APP_ANTIVIRUS_ENABLE") == "true"
	address := os.Getenv("APP_ANTIVIRUS_ADDRESS")
	timeoutStr := os.Getenv("APP_ANTIVIRUS_TIMEOUT")
	if len(timeoutStr) == 0 {
		timeoutStr = "10"
	}
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		return nil, errors.New("antivirus timeout not found")
	}
	network := os.Getenv("APP_ANTIVIRUS_NETWORK")

	if useAntivirus && len(address) == 0 {
		return nil, errors.New("antivirus address not found")
	}

	return &antivirusConfig{
		useAntivirus: useAntivirus,
		address:      address,
		timeout:      time.Duration(timeout) * time.Second,
		network:      network,
	}, nil
}

func (c *antivirusConfig) UseAntivirus() bool {
	return c.useAntivirus
}

func (c *antivirusConfig) Address() string {
	return c.address
}

func (c *antivirusConfig) Timeout() time.Duration {
	return c.timeout
}

func (c *antivirusConfig) Network() string {
	return c.network
}
