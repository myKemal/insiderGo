package main

import (
	"github.com/myKemal/mongoApi/app"
	_ "github.com/swaggo/http-swagger"
)

// @title insiderGo Project
// @version 1.0
// @description This is a MongoAPI project for managing messages.
// @host localhost:8090
// @BasePath /
func main() {
	application := app.NewApp()
	application.Run()
}
