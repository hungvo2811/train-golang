package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"test/src/dao"
	"test/src/models/payload"
	"test/src/models/query"
	"test/src/models/response"
)

func CreateDepartment(ctx context.Context, payload payload.DepartmentPayLoad) (res response.DepartmentResponse, err error) {
	department := payload.ConvertToBSON()
	if dao.GetDepartmentByName(department.Name) {
		err = errors.New("name already taken")
		return res, err
	}
	result, err := dao.CreateDepartment(ctx, department)
	if err != nil {
		err = errors.New("create department failed")
		return res, err
	}

	res = response.DepartmentResponse{
		ID:      result.ID.Hex(),
		Name:    result.Name,
		Address: result.Address,
	}

	return res, err
}

func UpdateDepartment(ctx context.Context, id primitive.ObjectID, payload payload.DepartmentPayLoad) (res response.DepartmentResponse, err error) {
	department := payload.ConvertToBSON()
	result, err := dao.UpdateDepartment(ctx, id, department)
	if err != nil {
		return res, errors.New("Update department failed")
	}

	res = response.DepartmentResponse{
		ID:      result.ID.Hex(),
		Name:    result.Name,
		Address: result.Address,
	}

	return res, err
}

func DeleteDepartment(ctx context.Context, id primitive.ObjectID) error {
	err := dao.DeleteDepartment(ctx, id)
	if err != nil {
		err = errors.New("delete department failed")
		return err
	}

	return nil
}

func GetDepartment(ctx context.Context, id primitive.ObjectID) (res response.DepartmentResponse, err error) {
	result, err := dao.GetDepartment(ctx, id)
	if err != nil {
		return res, errors.New("department not found")
	}

	res = response.DepartmentResponse{
		ID:      result.ID.Hex(),
		Name:    result.Name,
		Address: result.Address,
	}

	return res, err
}

func GetDepartments(ctx context.Context, query query.DepartmentQuery) (res []response.DepartmentResponse, err error) {
	result, err := dao.GetDepartments(ctx, query)
	if err != nil {
		return res, errors.New("get departments failed")
	}

	for _, value := range result {

		res = append(res, response.DepartmentResponse{
			ID:      value.ID.Hex(),
			Name:    value.Name,
			Address: value.Address,
		})
	}

	return res, err
}
