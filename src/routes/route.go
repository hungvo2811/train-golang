package routes

import "github.com/labstack/echo/v4"

func Route(e *echo.Echo) {
	StaffRouter(e)
	DepartmentRouter(e)
	AuthRouter(e)
}
