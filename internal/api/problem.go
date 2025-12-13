package api

import (
	"ShawnOJ/internal"
	"ShawnOJ/internal/api/params"
	"ShawnOJ/internal/api/responses"
	"ShawnOJ/internal/service"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AdminProblemSetHandler(c *gin.Context) {
	var ProblemSetParam params.AdminProblemSetParam
	if err := c.ShouldBind(&ProblemSetParam); err != nil {
		zap.L().Error("参数绑定错误：", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	id, ok := c.Get("userid")
	if !ok {
		zap.L().Error("id为空")
		responses.ResponseError(c, responses.CodeServerBusy)
	}
	adminId := id.(int64)
	ProblemSetParam.CreatorID = adminId

	if err := service.AdminProblemSetService(ProblemSetParam); err != nil {
		zap.L().Error("题目上传失败", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	responses.ResponseSuccess(c, nil)
}

func AdminProblemUpdateHandler(c *gin.Context) {
	var ProblemSetParam params.AdminProblemSetParam
	if err := c.ShouldBind(&ProblemSetParam); err != nil {
		zap.L().Error("参数绑定错误：", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}
}

func NormalGetProblemByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	roleStr, ok := c.Get("role")
	if !ok {
		zap.L().Error("参数错误", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}
	role := roleStr.(string)
	info, err := service.NormalGetProblemByIdService(id, role)
	if err != nil {
		if errors.Is(err, internal.ErrPermissionDenied) {
			zap.L().Error("查询错误：", zap.Error(err))
			responses.ResponseError(c, responses.CodeNotAdmin)
			return
		}
		zap.L().Error("查询错误：", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	responses.ResponseSuccess(c, info)
}

func NormalGetProblemBySlugHandler(c *gin.Context) {
	slug := c.Param("slug")

	roleStr, ok := c.Get("role")
	if !ok {
		zap.L().Error("无法获取用户身份")
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}
	role := roleStr.(string)
	info, err := service.NormalGetProblemBySlugService(slug, role)
	if err != nil {
		if errors.Is(err, internal.ErrPermissionDenied) {
			zap.L().Error("查询错误：", zap.Error(err))
			responses.ResponseError(c, responses.CodeNotAdmin)
			return
		}
		zap.L().Error("查询错误：", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	responses.ResponseSuccess(c, info)
}

func NormalGetProblemListHandler(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}
	sizeStr := c.DefaultQuery("size", "5")
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	roleStr, ok := c.Get("role")
	if !ok {
		zap.L().Error("无法获取用户身份")
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}
	role := roleStr.(string)
	difficultyStr := c.Param("difficulty")
	difficulty, err := strconv.ParseInt(difficultyStr, 10, 64)
	if difficulty > 3 || difficulty < 0 {
		zap.L().Error("参数错误", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	problemList, err := service.NormalGetProblemListService(role, page, size, difficulty)
	if err != nil {
		zap.L().Error("获取失败", zap.Error(err))
		responses.ResponseError(c, responses.CodeServerBusy)
		return
	}

	responses.ResponseSuccess(c, problemList)
}
