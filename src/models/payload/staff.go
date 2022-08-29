package payload

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"test/src/models"
)

type StaffPayLoad struct {
	ID           string `json:"_id"`
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password,omitempty" form:"password"`
	DepartmentID string `json:"department"  form:"department"`
}

type StaffLoginPayLoad struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
}

func (payload StaffPayLoad) ValidateStaff() error {
	return validation.ValidateStruct(&payload,
		validation.Field(
			&payload.Name,
			validation.Required.Error("Name is required"),
			is.Alpha.Error("Name is alphabet"),
		),
		validation.Field(
			&payload.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("Invalid email"),
		),
		validation.Field(
			&payload.Password,
			validation.Required.Error("Password is required"),
			validation.Length(6, 20).Error("Password length must be 6 to 20"),
		),
	)
}

func (payload StaffPayLoad) ValidateAuth() error {
	return validation.ValidateStruct(&payload,
		validation.Field(
			&payload.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("Invalid email"),
		),
		validation.Field(
			&payload.Password,
			validation.Required.Error("Password is required"),
			validation.Length(6, 20).Error("Password length must be 6 to 20"),
		),
	)
}

func (payload StaffPayLoad) ConvertToBSON() models.Staff {
	departmentID, _ := primitive.ObjectIDFromHex(payload.DepartmentID)
	result := models.Staff{
		ID:           primitive.NewObjectID(),
		Name:         payload.Name,
		Email:        payload.Email,
		Password:     payload.Password,
		DepartmentID: departmentID,
	}
	return result
}
