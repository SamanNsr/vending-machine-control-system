package infrastructure

import (
	"context"
	"go.uber.org/zap"

	"github.com/samannsr/vending-machine-control-system/pkg/config"
	echoHttp "github.com/samannsr/vending-machine-control-system/pkg/http/echo"
	"github.com/samannsr/vending-machine-control-system/pkg/logger"
)

type IContainer struct {
	Config         *config.Config
	Logger         *zap.Logger
	EchoHttpServer echoHttp.ServerInterface
}

func NewIC(ctx context.Context) (*IContainer, func(), error) {
	var downFns []func()
	down := func() {
		for _, df := range downFns {
			df()
		}
	}

	echoServerConfig := &echoHttp.ServerConfig{
		Port:     config.BaseConfig.Http.Port,
		BasePath: "/api/v1",
		IsDev:    config.IsDevEnv(),
	}
	echoServer := echoHttp.NewServer(echoServerConfig)
	echoServer.SetupDefaultMiddlewares()
	downFns = append(downFns, func() {
		_ = echoServer.GracefulShutdown(ctx)
	})

	ic := &IContainer{
		Config:         config.BaseConfig,
		Logger:         logger.Zap,
		EchoHttpServer: echoServer,
	}

	return ic, down, nil
}
