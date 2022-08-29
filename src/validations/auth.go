package validations

import (
	"github.com/labstack/echo/v4"
	"test/src/models/payload"
)

func Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload payload.StaffPayLoad
		)

		c.Bind(&payload)

		err := payload.ValidateAuth()

		if err != nil {
			return c.JSON(400, err)
		}

		c.Set("payload", payload)
		return next(c)
	}
}
