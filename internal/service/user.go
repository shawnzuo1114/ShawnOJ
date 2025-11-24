package service

import (
	"ShawnOJ/internal"
	"ShawnOJ/internal/api/params"
	"ShawnOJ/internal/dao"
	"ShawnOJ/internal/models"
	"ShawnOJ/utils/bcrypt"
	"ShawnOJ/utils/jwt"
	"ShawnOJ/utils/mail"
	"ShawnOJ/utils/snowflake"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
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

func UserLoginService(lp params.LoginParam) (token string, err error) {
	oPassword := lp.Password
	user := &models.User{
		Username: lp.Username,
		Password: lp.Password,
	}
	if err = dao.QueryUser(user.Username); err != nil {
		return
	}
	hashPassword, err := dao.SelectUserPassword(user)
	if err != nil {
		return "", err
	}
	err = bcrypt.CheckPasswordHash(oPassword, hashPassword)
	if err != nil {
		return
	}
	token, err = jwt.GenToken(user.UserID, user.Username)
	return token, err
}

func UserInfoService(ip *params.InfoParam) (err error) {
	u := &models.User{
		UserID: ip.UserID,
	}
	if err = dao.SelectUserInfoByID(u); err != nil {
		return
	}
	ip.Username = u.Username
	ip.UserID = u.UserID
	ip.Role = u.Role
	ip.CreatedAt = u.CreatedAt
	ip.UpdatedAt = u.UpdatedAt
	return
}

func UserPasswordService(pp *params.PasswordParam) (err error) {
	oPassword := pp.OldPassword
	hashPassword, err := dao.CheckUserPassword(pp.UserID)
	if err != nil {
		return
	}
	if err = bcrypt.CheckPasswordHash(oPassword, hashPassword); err != nil {
		return
	}
	newPassword, err := bcrypt.GetPasswordHash(pp.NewPassword)
	if err != nil {
		return
	}
	user := &models.User{
		UserID:   pp.UserID,
		Password: newPassword,
	}
	if err = dao.UpdatePassword(user); err != nil {
		return
	}
	return
}

func UserLogoutService(token string) (err error) {
	fmt.Println(token)
	return nil
}

func UserCaptchaService(email string, c *gin.Context) (err error) {
	flag, err := dao.GetCodeTTL(email, c)
	if err != nil {
		return
	}
	if flag {
		return internal.ErrVerifyExists
	}

	code, err := mail.SendVerifyCode(email)
	if err != nil {
		return
	}

	if err = dao.SetVerifyCode(email, code, c); err != nil {
		return
	}

	return
}

func UserVerifyService(email string, code int, c *gin.Context) (err error) {
	value, err := dao.GetVerifyCode(email, c)
	if err != nil {
		return
	}
	rCode, _ := strconv.Atoi(value)
	if rCode != code {
		return internal.ErrVerifyCodeWrong
	}
	if err = dao.DeleteVerifyCode(email, c); err != nil {
		return
	}
	return nil
}
