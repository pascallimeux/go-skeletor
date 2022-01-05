package main

import (
	"fmt"

	"example.com/skeleton/config"
	"example.com/skeleton/internal/examples"
	"example.com/skeleton/log"
)


func main() {
	appConfig := config.AppConfig
	logger := log.New(appConfig.LogLevel)
	fmt.Println(appConfig)
	logger.Trace("message de trace")
	logger.Debug("message de debug")
	logger.Info("message d'info")
	logger.Warn("message de warning")
	logger.Error("message d'erreur")
	logger.Fatal("message fatal")
	//logger.Panic("message panic")
	fmt.Println(examples.Hello())
}
