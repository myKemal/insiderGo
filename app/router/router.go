package router

import (
	"github.com/gorilla/mux"
	_ "github.com/myKemal/mongoApi/app/docs"
	"github.com/myKemal/mongoApi/app/handler"
	"github.com/myKemal/mongoApi/app/middleware"
	"github.com/myKemal/mongoApi/app/repository"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func InitializeRoutes(mongoRepo repository.MongoRepository, tempMemory repository.TempMemoryRepo) http.Handler {
	router := mux.NewRouter()

	h := handler.NewHandler(mongoRepo, tempMemory)

	router.HandleFunc("/generate-token", h.GenerateToken).Methods("POST")

	router.HandleFunc("/unsent-messages", h.FetchUnsentMessages).Methods("GET").
		Handler(middleware.AuthMiddleware(http.HandlerFunc(h.FetchUnsentMessages)))

	router.HandleFunc("/temp/getList", h.GetKeys).Methods("POST").
		Handler(middleware.AuthMiddleware(http.HandlerFunc(h.GetKeys)))

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	corsMiddleware := middleware.CORSHandler()
	return corsMiddleware(router)
}
