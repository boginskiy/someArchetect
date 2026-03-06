package config

import (
	"errors"
	"net"
	"os"

	"github.com/joho/godotenv"
)

const (
	httpHostEnvName = "HOST"
	grpcPortEnvName = "PORT"
)

type EnvConf struct {
	Host string
	Port string
}

func NewEnvConf(pathToEnv string) (*EnvConf, error) {
	conf := &EnvConf{}

	// Default
	if pathToEnv == "" {
		conf.Host = HostName
		conf.Port = PortName
		return conf, nil
	}

	conf.load(pathToEnv)

	// Host
	conf.Host = os.Getenv(httpHostEnvName)
	if len(conf.Host) == 0 {
		return nil, errors.New("http host not found")
	}

	// Port
	conf.Port = os.Getenv(grpcPortEnvName)
	if len(conf.Port) == 0 {
		return nil, errors.New("http port not found")
	}

	return conf, nil
}

func (c *EnvConf) load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *EnvConf) GetAddress() string {
	return net.JoinHostPort(c.Host, c.Port)
}
