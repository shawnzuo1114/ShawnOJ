package api

import (
	"ShawnOJ/internal"
	"ShawnOJ/internal/api/params"
	"ShawnOJ/internal/api/responses"
	"ShawnOJ/internal/service"
	utils "ShawnOJ/utils/regexp"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserRegisterHandler(c *gin.Context) {
	var registerParam params.RegisterParam

	// 绑定结构体
	if err := c.ShouldBindJSON(&registerParam); err != nil {
		zap.L().Error("传参错误", zap.Error(err))
		responses.ResponseError(c, responses.CodeInvaildParam)
		return
	}

	// 校验密码是否相同
	if registerParam.Password != registerParam.RePassword {
		zap.L().Error("两次密码不一致")
		responses.ResponseError(c, responses.CodeInvaildParam)
		return
	}

	// 正则表达式校验密码格式是否正确
	if err := utils.Match(utils.PasswordPattern, registerParam.Password); err != nil {
		zap.L().Error("密码错误：", zap.Error(err))
		responses.ResponseError(c, responses.CodeInvaildPassword)
		return
	}

	// 送入service层
	if err := service.UserRegisterService(registerParam); err != nil {
		// 判断用户是否存在
		if errors.Is(err, internal.ErrUserExists) {
			zap.L().Error("用户已存在", zap.Error(err))
			responses.ResponseError(c, responses.CodeUserExist)
			return
		}
		zap.L().Error("系统错误：", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	// 成功返回
	responses.ResponseSuccess(c, nil)
}
