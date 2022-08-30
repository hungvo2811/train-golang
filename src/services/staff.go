package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"test/src/dao"
	"test/src/models/payload"
	"test/src/models/query"
	"test/src/models/response"
)

func CreateStaff(ctx context.Context, payload payload.StaffPayLoad) (res response.StaffResponse, err error) {
	departmentID, _ := primitive.ObjectIDFromHex(payload.DepartmentID)
	getDepartment, err2 := GetDepartment(ctx, departmentID)
	if err2 != nil {
		err = errors.New("department Not Found")
		return res, err
	}
	payload.DepartmentID = getDepartment.ID
	staff := payload.ConvertToBSON()
	if dao.GetStaffByEmail(ctx, payload.Email) {
		err = errors.New("email Already Taken")
		return res, err
	}
	result, err := dao.CreateStaff(ctx, staff)
	if err != nil {
		err = errors.New("create Staff Failed")
		return res, err
	}

	res = response.StaffResponse{
		ID:       result.ID.Hex(),
		Name:     result.Name,
		Email:    result.Email,
		Password: "",
		Department: response.DepartmentResponse{
			ID:      getDepartment.ID,
			Name:    getDepartment.Name,
			Address: getDepartment.Address,
		},
	}

	return res, err
}

func UpdateStaff(ctx context.Context, id primitive.ObjectID, payload payload.StaffPayLoad) (res response.StaffResponse, err error) {
	staff := payload.ConvertToBSON()
	result, err := dao.UpdateStaff(ctx, id, staff)

	if err != nil {
		return res, errors.New("update Staff Failed")
	}

	res = response.StaffResponse{
		ID:       result.ID.Hex(),
		Name:     result.Name,
		Email:    result.Email,
		Password: "",
		Department: response.DepartmentResponse{
			ID: result.DepartmentID.Hex(),
		},
	}

	return res, err
}

func GetStaff(ctx context.Context, id primitive.ObjectID) (res response.StaffResponse, err error) {
	result, err := dao.GetStaff(ctx, id)

	if err != nil {
		return res, errors.New("staff Not Found")
	}

	res = response.StaffResponse{
		ID:    result.ID.Hex(),
		Name:  result.Name,
		Email: result.Email,
		Department: response.DepartmentResponse{
			ID: result.ID.Hex(),
		},
	}

	return res, err
}

func GetStaffs(ctx context.Context, query query.StaffQuery) (res []response.StaffResponse, err error) {
	rs, err := dao.GetStaffs(ctx, query)
	if err != nil {
		return res, errors.New("get Staffs Failed")
	}

	for _, value := range rs {
		department, _ := GetDepartment(ctx, value.DepartmentID)

		res = append(res, response.StaffResponse{
			ID:         value.ID.Hex(),
			Name:       value.Name,
			Email:      value.Email,
			Password:   value.Password,
			Department: department,
		})
	}

	return res, err
}

func DeleteStaff(ctx context.Context, id primitive.ObjectID) error {
	err := dao.DeleteStaff(ctx, id)

	if err != nil {
		err = errors.New("delete Staff Failed")
		return err
	}
	return nil
}
