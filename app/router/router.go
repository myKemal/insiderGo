package router

import (
	"github.com/gorilla/mux"
	_ "github.com/myKemal/insiderGo/app/docs"
	"github.com/myKemal/insiderGo/app/handler"
	"github.com/myKemal/insiderGo/app/middleware"
	"github.com/myKemal/insiderGo/app/repository"
	"github.com/myKemal/insiderGo/app/services"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func InitializeRoutes(mongoRepo repository.MongoRepository, tempMemory repository.TempMemoryRepository, taskService *services.PeriodicTaskService) http.Handler {
	router := mux.NewRouter()

	//tek handler iki dependency
	h := handler.NewHandler(mongoRepo, tempMemory)

	router.HandleFunc("/", handler.HomeHandler)

	router.HandleFunc("/generate-token", h.GenerateToken).Methods("POST")

	router.HandleFunc("/unsent-messages", h.FetchUnsentMessages).Methods("GET").
		Handler(middleware.AuthMiddleware(http.HandlerFunc(h.FetchUnsentMessages)))

	router.HandleFunc("/temp/getList", h.GetList).Methods("POST").
		Handler(middleware.AuthMiddleware(http.HandlerFunc(h.GetList)))

	router.HandleFunc("/temp/getAllList", h.GetAllList).Methods("GET").
		Handler(middleware.AuthMiddleware(http.HandlerFunc(h.GetAllList)))

	router.HandleFunc("/instance/unsent-messages", handler.GetUnsentMessagesHandler).Methods("GET").
		Handler(middleware.AuthMiddleware(http.HandlerFunc(handler.GetUnsentMessagesHandler)))

	// ayrı handler tek dependency
	taskHandler := handler.NewPeriodicTaskHandler(taskService)
	router.HandleFunc("/task/start", taskHandler.StartPeriodicTaskService).Methods("POST")
	router.HandleFunc("/task/stop", taskHandler.StopPeriodicTaskService).Methods("POST")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	corsMiddleware := middleware.CORSHandler()
	return corsMiddleware(router)
}
