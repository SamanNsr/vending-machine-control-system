package main

import (
	"github.com/samannsr/vending-machine-control-system/app"
	"github.com/samannsr/vending-machine-control-system/pkg/logger"
)

func main() {
	err := app.New().Run()
	if err != nil {
		logger.Zap.Sugar().Fatal(err)
	}
}
