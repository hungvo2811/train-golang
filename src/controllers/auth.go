package controllers

import (
	"github.com/labstack/echo/v4"
	"test/src/middleware"
	"test/src/models/payload"
	"test/src/services"
	"test/src/util"
)

func Login(c echo.Context) error {
	ctx := c.Request().Context()
	var payload, _ = c.Get("payload").(payload.StaffLoginPayLoad)
	err := c.Bind(&payload)
	if err != nil {
		return err
	}
	err2 := services.Login(ctx, payload)
	if err2 != nil {
		return util.Response400(c, nil, err.Error())
	}
	token := middleware.CreateToken()
	return util.ResponseLogin(c, token, "Login Successful")
}
