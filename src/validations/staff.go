package validations

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"test/src/models/payload"
	"test/src/models/query"
)

func CreateStaff(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload payload.StaffPayLoad
		)

		c.Bind(&payload)

		err := payload.ValidateStaff()

		if err != nil {
			return c.JSON(400, err)
		}

		c.Set("payload", payload)
		return next(c)
	}
}

func UpdateStaff(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload payload.StaffPayLoad
		)

		c.Bind(&payload)

		err := payload.ValidateStaff()

		if err != nil {
			return c.JSON(400, err)
		}

		c.Set("payload", payload)
		return next(c)
	}
}

func GetStaffs(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			query query.StaffQuery
		)

		err := c.Bind(&query)
		if err != nil {
			c.JSON(400, errors.New("Invalid input"))
		}

		err = query.ValidateStaffQuery()
		if err != nil {
			c.JSON(400, errors.New("Invalid input"))
		}

		c.Set("query", query)
		return next(c)
	}
}

func ValidateObjectID(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err)
	}

	return objectID, err
}

func ValidateStaffID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id           = c.Param("id")
			staffID, err = ValidateObjectID(id)
		)

		if err != nil {
			return c.JSON(400, errors.New("Staff ID is not valid"))
		}

		c.Set("staffID", staffID)
		return next(c)
	}
}
