package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"test/src/database"
	"test/src/routes"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	database.ConnectDB()
	routes.Route(e)
	e.Logger.Fatal(e.Start(":8080"))
}
