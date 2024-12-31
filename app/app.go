package app

import (
	"fmt"
	"github.com/myKemal/mongoApi/app/initialize"
	"log"
	"net/http"
	"os"

	"github.com/myKemal/mongoApi/app/router"
)

type App struct {
	Router http.Handler
}

func NewApp() *App {

	mongoRepo, err := initialize.Mongo()
	if err != nil {
		return nil
	}

	tempRepository, err := initialize.Temp()
	if err != nil {
		return nil
	}

	myRouter := router.InitializeRoutes(mongoRepo, tempRepository)

	return &App{
		Router: myRouter,
	}
}

func (a *App) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}
