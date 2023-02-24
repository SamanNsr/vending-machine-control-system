package app

import (
	"context"
	vendingMachineConfigurator "github.com/samannsr/vending-machine-control-system/internal/vending_machine/configurator"
	"os"
	"os/signal"
	"syscall"

	"github.com/samannsr/vending-machine-control-system/pkg/infrastructure"
)

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ic, infraDown, err := infrastructure.NewIC(ctx)
	if err != nil {
		return err
	}
	defer infraDown()

	me := configureModule(ctx, ic)
	if me != nil {
		return me
	}

	var serverError error
	go func() {
		if err := ic.EchoHttpServer.RunServer(ctx, nil); err != nil {
			ic.Logger.Sugar().Errorf("(s.RunEchoServer) err: {%v}", err)
			serverError = err
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		ic.Logger.Sugar().Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		ic.Logger.Sugar().Errorf("ctx.Done: %v", done)
	}

	ic.Logger.Sugar().Info("Server Exited Properly")
	return serverError
}

func configureModule(ctx context.Context, ic *infrastructure.IContainer) error {
	err := vendingMachineConfigurator.NewConfigurator(ic).Configure(ctx)
	if err != nil {
		return err
	}

	return nil
}
