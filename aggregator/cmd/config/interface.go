package config

type Config interface {
	EnvConfig
	GetConnToDB() string
}

type EnvConfig interface {
	GetAddress() string
}
