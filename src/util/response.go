package util

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response map[string]interface{}

func generateResponse(data interface{}, message string) Response {
	return Response{
		"data":    data,
		"message": message,
	}
}

func generateResponseLogin(data interface{}, message string) Response {
	return Response{
		"access_token": data,
		"message":      message,
	}
}

func Response200(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Success!"
	}
	return c.JSON(http.StatusOK, generateResponse(data, message))
}

func ResponseLogin(c echo.Context, token interface{}, message string) error {
	if message == "" {
		message = "Success!"
	}
	return c.JSON(http.StatusOK, generateResponseLogin(token, message))
}

func Response400(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Data invalid"
	}
	return c.JSON(http.StatusBadRequest, generateResponse(data, message))
}

func Response404(c echo.Context, data interface{}, message string) error {

	if message == "" {
		message = "Data not found"
	}
	return c.JSON(http.StatusNotFound, generateResponse(data, message))
}
