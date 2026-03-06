package config

type Config interface {
	EnvConfig
}

type EnvConfig interface {
	GetAddress() string
}
