package services

import (
	"context"
	"errors"
	"test/src/dao"
	"test/src/models/payload"
)

func Login(ctx context.Context, payload payload.StaffLoginPayLoad) error {
	err := dao.GetStaffByEmailAndPassword(ctx, payload.Email, payload.Password)
	if err != true {
		err2 := errors.New("login Failed")
		return err2
	}
	return nil
}
