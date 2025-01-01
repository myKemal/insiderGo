package main

import (
	"github.com/myKemal/insiderGo/app"
	_ "github.com/swaggo/http-swagger"
	"os"
	"os/signal"
	"syscall"
)

// @title insiderGo Project
// @version 1.0
// @description This is a insiderGo task project for managing messages.
// @host localhost:8090
// @BasePath /
func main() {
	application := app.NewApp()

	stop := make(chan os.Signal, 1) //dışardan kapanırsa
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		application.Run()
	}()

	<-stop
	application.Stop()

}
