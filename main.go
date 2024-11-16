package main

import (
	"logging-api/config"
	"logging-api/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Database
	db := config.ConnectDB()

	// Routes
	loggingHandler := handlers.NewLoggingHandler(db)
	e.POST("/logging", handlers.Index)
	e.POST("/logging/add", loggingHandler.Create)
	e.POST("/logging/update", loggingHandler.Update)
	e.POST("/logging/list", loggingHandler.GetAll)
	e.POST("/logging/delete", loggingHandler.Delete)


	e.Logger.Fatal(e.Start("localhost:5000"))
}