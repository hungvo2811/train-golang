package routes

import (
	"github.com/labstack/echo/v4"
	"test/src/controllers"
	"test/src/validations"
)

func AuthRouter(e *echo.Echo) {
	departments := e.Group("")

	departments.POST("/login", controllers.Login, validations.Login)
}
