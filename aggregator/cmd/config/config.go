package config

type Conf struct {
	EnvConf EnvConfig
}

func NewConf(envConf EnvConfig) *Conf {
	return &Conf{
		EnvConf: envConf,
	}
}

func (c *Conf) GetAddress() string {
	return c.EnvConf.GetAddress()
}
