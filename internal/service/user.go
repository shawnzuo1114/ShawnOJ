package service

import (
	"ShawnOJ/internal/api/params"
	"ShawnOJ/internal/dao"
	"ShawnOJ/internal/models"
	"ShawnOJ/utils/bcrypt"
	"ShawnOJ/utils/snowflake"
)

func UserRegisterService(rp params.RegisterParam) (err error) {
	if err = dao.CheckUser(rp.Username); err != nil {
		return
	}
	oPassword := rp.Password
	password, err := bcrypt.GetPasswordHash(oPassword)
	if err != nil {
		return
	}
	user := &models.User{
		UserID:   snowflake.GenID(),
		Username: rp.Username,
		Password: password,
	}
	if err = dao.InsertUser(user); err != nil {
		return
	}
	return nil
}
