package payload

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"test/src/models"
)

type DepartmentPayLoad struct {
	ID      string `json:"_id" form:"_id"`
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}

func (payload DepartmentPayLoad) ConvertToBSON() models.Department {
	result := models.Department{
		ID:      primitive.NewObjectID(),
		Name:    payload.Name,
		Address: payload.Address,
	}
	return result
}

func (payload DepartmentPayLoad) ValidateCreateDepartment() error {
	return validation.ValidateStruct(&payload,
		validation.Field(
			&payload.Name,
			validation.Required.Error("Name is required"),
			validation.Length(3, 30).Error("Name is length: 3 -> 30"),
			is.Alpha.Error("Name is alphabet"),
		),
		validation.Field(
			&payload.Address,
			validation.Required.Error("Address is required"),
			validation.Length(10, 300).Error("Name is length: 3 -> 30"),
		),
	)
}
