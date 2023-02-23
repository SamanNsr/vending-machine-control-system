package echoHttp

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/samannsr/vending-machine-control-system/pkg/constant"
	echoErrorHandler "github.com/samannsr/vending-machine-control-system/pkg/http/echo/handlers/error_handler"
	"github.com/samannsr/vending-machine-control-system/pkg/logger"
	loggerConstant "github.com/samannsr/vending-machine-control-system/pkg/logger/constant"
)

type ServerConfig struct {
	Port     int
	BasePath string
	IsDev    bool
}

type Server struct {
	echo   *echo.Echo
	config *ServerConfig
}

type ServerInterface interface {
	RunServer(ctx context.Context, configEcho func(echoServer *echo.Echo)) error
	GracefulShutdown(ctx context.Context) error
	GetEchoInstance() *echo.Echo
	SetupDefaultMiddlewares()
	AddMiddlewares(middlewares ...echo.MiddlewareFunc)
	GetBasePath() string
}

func NewServer(config *ServerConfig) *Server {
	return &Server{echo: echo.New(), config: config}
}

func (s *Server) RunServer(ctx context.Context, configEcho func(echo *echo.Echo)) error {
	if configEcho != nil {
		configEcho(s.echo)
	}

	go func() {
		<-ctx.Done()
		logger.Zap.Sugar().Infof("Http server is shutting down PORT: %d", s.config.Port)
		if err := s.GracefulShutdown(ctx); err != nil {
			logger.Zap.Sugar().Warnf("(Shutdown) err: {%v}", err)
		}
	}()

	logger.Zap.Sugar().Infof("[echoServer.RunHttpServer] Echo server is listening on: %d", s.config.Port)
	return s.echo.Start(fmt.Sprintf(":%d", s.config.Port))
}

func (s *Server) AddMiddlewares(middlewares ...echo.MiddlewareFunc) {
	if len(middlewares) > 0 {
		s.echo.Use(middlewares...)
	}
}

func (s *Server) GracefulShutdown(ctx context.Context) error {
	err := s.echo.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) SetupDefaultMiddlewares() {
	s.echo.HideBanner = false
	s.echo.Pre(middleware.RemoveTrailingSlash())
	s.echo.HTTPErrorHandler = echoErrorHandler.ErrorHandler

	s.echo.Use(middleware.Recover())

	s.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogError:     false,
		LogMethod:    true,
		LogRequestID: true,
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			t := time.Now()
			logger.Zap.Info(
				"Incoming Request",
				zap.String(loggerConstant.TYPE, loggerConstant.HTTP),
				zap.String(loggerConstant.METHOD, v.Method),
				zap.String(loggerConstant.REQUEST_ID, v.RequestID),
				zap.String(loggerConstant.URI, v.URI),
				zap.String(loggerConstant.STATUS, http.StatusText(v.Status)),
				zap.Duration(loggerConstant.LATENCY, v.Latency),
				zap.Time(loggerConstant.TIME, t),
			)
			return nil
		},
	}))

	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: constant.EchoGzipLevel,
		// Skipper: func(c echo.Context) bool {
		//	return strings.Contains(c.Request().URL.Path, "swagger")
		// },
	}))
}

func (s *Server) GetEchoInstance() *echo.Echo {
	return s.echo
}

func (s *Server) GetBasePath() string {
	return s.config.BasePath
}
