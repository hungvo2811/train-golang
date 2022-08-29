package validations

import (
	"errors"
	"github.com/labstack/echo/v4"
	"test/src/models/payload"
	"test/src/models/query"
)

func CreateDepartment(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload payload.DepartmentPayLoad
		)

		c.Bind(&payload)

		err := payload.ValidateCreateDepartment()

		if err != nil {
			return c.JSON(400, err)
		}

		c.Set("payload", payload)
		return next(c)
	}
}

func UpdateDepartment(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload payload.DepartmentPayLoad
		)

		c.Bind(&payload)

		err := payload.ValidateCreateDepartment()

		if err != nil {
			return c.JSON(400, err)
		}

		c.Set("payload", payload)
		return next(c)
	}
}

func ValidateDepartmentID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id                = c.Param("id")
			departmentID, err = ValidateObjectID(id)
		)

		if err != nil {
			return c.JSON(400, errors.New("Department ID is not valid"))
		}

		c.Set("departmentID", departmentID)
		return next(c)
	}
}

func GetDepartments(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			query query.DepartmentQuery
		)

		err := c.Bind(&query)
		if err != nil {
			c.JSON(400, errors.New("Invalid input"))
		}

		err = query.ValidateDepartmentQuery()
		if err != nil {
			c.JSON(400, errors.New("Invalid input"))
		}

		c.Set("query", query)
		return next(c)
	}
}
