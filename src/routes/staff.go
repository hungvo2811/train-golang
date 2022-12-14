package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"test/src/controllers"
	"test/src/validations"
)

func StaffRouter(e *echo.Echo) {
	staffs := e.Group("/staffs")
	staffs.Use(middleware.JWT([]byte("secret")))
	staffs.POST("", controllers.CreateStaff, validations.CreateStaff)
	staffs.PUT("/:id", controllers.UpdateStaff, validations.ValidateStaffID, validations.UpdateStaff)
	staffs.DELETE("/:id", controllers.DeleteStaff, validations.ValidateStaffID)
	staffs.GET("/:id", controllers.GetStaff, validations.ValidateStaffID)
	staffs.GET("", controllers.GetStaffs, validations.GetStaffs)

}
