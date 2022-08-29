package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"test/src/models/payload"
	"test/src/models/query"
	"test/src/services"
	"test/src/util"
)

func CreateStaff(c echo.Context) error {
	ctx := c.Request().Context()
	var payload, ok = c.Get("payload").(payload.StaffPayLoad)

	if !ok {
		return util.Response400(c, nil, "Bad Request")
	}

	result, err := services.CreateStaff(ctx, payload)
	if err != nil {
		return util.Response400(c, nil, "Create Staff Failed")
	}

	return util.Response200(c, result, "Create Staff Successful")
}

func UpdateStaff(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))
	var payload, ok = c.Get("payload").(payload.StaffPayLoad)
	if !ok {
		return util.Response400(c, nil, "Bad Request")
	}

	result, err := services.UpdateStaff(ctx, ID, payload)
	if err != nil {
		return util.Response400(c, nil, "Update Staff Failed")
	}

	return util.Response200(c, result, "Update Staff Successful")
}

func GetStaff(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))

	result, err := services.GetStaff(ctx, ID)
	if err != nil {
		return util.Response400(c, nil, "Get Staff Failed")
	}

	return util.Response200(c, result, "Get Staff Successful")
}

func GetStaffs(c echo.Context) error {
	ctx := c.Request().Context()
	query, ok := c.Get("query").(query.StaffQuery)
	if !ok {
		return c.JSON(400, "Bad Request")
	}

	result, err := services.GetStaffs(ctx, query)
	if err != nil {
		return util.Response400(c, nil, "Get Staffs Failed")
	}

	return util.Response200(c, result, "Get Staffs Successful")
}

func DeleteStaff(c echo.Context) error {
	ctx := c.Request().Context()
	var ID, _ = primitive.ObjectIDFromHex(c.Param("id"))

	err := services.DeleteStaff(ctx, ID)
	if err != nil {
		return util.Response400(c, nil, "Get Staffs Failed")
	}

	return util.Response200(c, nil, "Delete Staffs Successful")
}
