package api

import (
	"ShawnOJ/internal"
	"ShawnOJ/internal/api/params"
	"ShawnOJ/internal/api/responses"
	"ShawnOJ/internal/service"
	"ShawnOJ/utils/jwt"
	"ShawnOJ/utils/regexp"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UserRegisterHandler 用户注册
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

// UserLoginHandler 用户登录
func UserLoginHandler(c *gin.Context) {
	var loginParam params.LoginParam

	// 绑定结构体
	if err := c.ShouldBindJSON(&loginParam); err != nil {
		zap.L().Error("传参错误:", zap.Error(err))
		responses.ResponseError(c, responses.CodeInvaildParam)
		return
	}

	// 送入service层处理逻辑
	token, err := service.UserLoginService(loginParam)
	if err != nil {
		zap.L().Error("登录错误：", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	responses.ResponseSuccess(c, token)
}

// UserInfoHandler 用户信息查询逻辑
func UserInfoHandler(c *gin.Context) {
	// JWT Claims 中的数字通常是 float64
	var uid int64
	uid, err := jwt.GetUserID(c)
	if err != nil {
		zap.L().Error("userid 类型错误", zap.Error(err))
		responses.ResponseError(c, responses.CodeInvalidID)
		return
	}

	// 查询信息
	var ip params.InfoParam
	ip.UserID = uid

	if err = service.UserInfoService(&ip); err != nil {
		zap.L().Error("信息获取失败：", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	responses.ResponseSuccess(c, ip)
}

// UserPasswordHandler 修改密码
func UserPasswordHandler(c *gin.Context) {
	var passwordParam params.PasswordParam
	uid, err := jwt.GetUserID(c)
	if err != nil {
		zap.L().Error("userid 类型错误", zap.Error(err))
		responses.ResponseError(c, responses.CodeInvalidID)
		return
	}
	passwordParam.UserID = uid
	if err = c.ShouldBindJSON(&passwordParam); err != nil {
		zap.L().Error("参数绑定错误", zap.Error(err))
		responses.ResponseError(c, responses.CodeInvaildParam)
		return
	}

	if err = service.UserPasswordService(&passwordParam); err != nil {
		zap.L().Error("修改失败", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}
	responses.ResponseSuccess(c, nil)
}

func UserLogoutHandler(c *gin.Context) {
	token, ok := c.Get("token")
	if token == nil || !ok {
		zap.L().Error("token为空")
		responses.ResponseError(c, responses.CodeInvalidAuth)
		return
	}
	if err := service.UserLogoutService(token.(string)); err != nil {
		zap.L().Error("逻辑错误：", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}
	return
}
