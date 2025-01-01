package app

import (
	"fmt"
	"github.com/myKemal/insiderGo/app/initialize"
	"github.com/myKemal/insiderGo/app/services"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/myKemal/insiderGo/app/router"
)

type App struct {
	Router              http.Handler
	PeriodicTaskService *services.PeriodicTaskService
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

	messageService := services.GetMessageService()

	err = initialize.Start(mongoRepo, tempRepository, messageService)
	if err != nil {
		log.Fatalf("Failed to start initialize: %v", err)
	}

	periodicWebhookService := services.NewPeriodicTaskService(messageService, tempRepository, services.NewWebhookService(), 2*time.Minute)

	myRouter := router.InitializeRoutes(mongoRepo, tempRepository, periodicWebhookService)

	_, err = periodicWebhookService.Start()
	if err != nil {
		return nil
	}

	return &App{
		Router:              myRouter,
		PeriodicTaskService: periodicWebhookService,
	}
}

func (a *App) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, a.Router))

}

func (a *App) Stop() {
	if a.PeriodicTaskService != nil {
		_, err := a.PeriodicTaskService.Stop()
		if err != nil {
			return
		}
	}
}
