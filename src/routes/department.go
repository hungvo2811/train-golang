package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"test/src/controllers"
	"test/src/validations"
)

func DepartmentRouter(e *echo.Echo) {
	departments := e.Group("/departments")
	departments.Use(middleware.JWT([]byte("secret")))
	departments.POST("", controllers.CreateDepartment, validations.CreateDepartment)
	departments.PUT("/:id", controllers.UpdateDepartment, validations.ValidateDepartmentID, validations.UpdateDepartment)
	departments.DELETE("/:id", controllers.DeleteDepartment, validations.ValidateDepartmentID)
	departments.GET("/:id", controllers.GetDepartment, validations.ValidateDepartmentID)
	departments.GET("", controllers.GetDepartments, validations.GetDepartments)

}
