package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"test/src/models/payload"
	"test/src/models/query"
	"test/src/services"
	"test/src/util"
)

func CreateDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	var payload, ok = c.Get("payload").(payload.DepartmentPayLoad)
	if !ok {
		return util.Response400(c, nil, "Bad Request")
	}

	result, err := services.CreateDepartment(ctx, payload)
	if err != nil {
		return util.Response400(c, nil, "Create Department Failed")
	}

	return util.Response200(c, result, "Create Department Successful")
}

func UpdateDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))
	var payload, ok = c.Get("payload").(payload.DepartmentPayLoad)
	if !ok {
		return util.Response400(c, nil, "Bad Request")
	}

	result, err := services.UpdateDepartment(ctx, ID, payload)
	if err != nil {
		return util.Response200(c, nil, "Update Department Failed")
	}

	return util.Response200(c, result, "Update Department Successful")
}

func DeleteDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))

	err := services.DeleteDepartment(ctx, ID)
	if err != nil {
		return util.Response400(c, nil, "Delete Department failed")
	}

	return util.Response200(c, nil, "Delete Department Successful")
}

func GetDepartment(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))

	result, err := services.GetDepartment(ctx, ID)
	if err != nil {
		return util.Response400(c, nil, "Get Department Failed")
	}

	return util.Response200(c, result, "Get Department Successful")
}

func GetDepartments(c echo.Context) error {
	ctx := c.Request().Context()
	query, ok := c.Get("query").(query.DepartmentQuery)
	if !ok {
		return util.Response400(c, nil, "Bad Request")
	}

	result, err := services.GetDepartments(ctx, query)
	if err != nil {
		return util.Response400(c, nil, "Get Departments Failed")
	}

	return util.Response200(c, result, "Get Departments Successful")
}
