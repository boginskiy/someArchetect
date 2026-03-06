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

func (c *Conf) GetConnToDB() string {
	return "postgres://username:userpassword@localhost:5432/testdb?sslmode=disable"
}
