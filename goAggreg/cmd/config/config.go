package config

import "context"

type Config interface {
	GetAddress() string
}

type Cfg struct {
}

func NewCfg(_ context.Context) (*Cfg, error) {
	return &Cfg{}, nil
}

func (c *Cfg) GetAddress() string {
	return "localhost:8080"
}
