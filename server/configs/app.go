package configs

type AppConfig struct {
	ListenPort    int
}

func (c *AppConfig) LoadAppConfig() {
	c.ListenPort = 9090
}
