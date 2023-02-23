package config

import (
	"github.com/samannsr/vending-machine-control-system/pkg/constant"
	"github.com/samannsr/vending-machine-control-system/pkg/env"
)

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
	return &Config{
		App: AppConfig{
			AppEnv:  env.New("APP_ENV", constant.AppEnvDev).AsString(),
			AppName: env.New("APP_NAME", constant.AppName).AsString(),
		},
		Http: HttpConfig{
			Port: env.New("HTTP_PORT", constant.HttpPort).AsInt(),
		},
	}
}

func IsDevEnv() bool {
	return BaseConfig.App.AppEnv == constant.AppEnvDev
}

func IsProdEnv() bool {
	return BaseConfig.App.AppEnv == constant.AppEnvProd
}

func IsTestEnv() bool {
	return BaseConfig.App.AppEnv == constant.AppEnvTest
}
