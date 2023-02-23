package config

var BaseConfig *Config

type Config struct {
	App  AppConfig
	Http HttpConfig
}

type AppConfig struct {
	AppEnv  string
	AppName string
}

type HttpConfig struct {
	Port int
}

func init() {
	BaseConfig = newConfig()
}

func newConfig() *Config {
	return &Config{}
}
