package routers

import (
	"ShawnOJ/internal/api"
	"ShawnOJ/internal/middleware"
	"ShawnOJ/utils/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	var r *gin.Engine
	if mode == "dev" {
		r = gin.Default()
		r.Use(logger.GinLogger(), logger.GinRecovery(true))
	} else {
		r = gin.New()
		r.Use(logger.GinLogger(), logger.GinRecovery(true))
	}

	r.Use(middleware.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
		})
	})

	apiGroup1 := r.Group("/api1/v1")

	apiGroup1.POST("user/register", api.UserRegisterHandler)
	apiGroup1.GET("user/login", api.UserLoginHandler)

	apiGroup2 := r.Group("/api2/v1")
	userGroup := apiGroup2.Group("user")
	userGroup.Use(middleware.JWTAuthMiddleware())
	userGroup.GET("me", api.UserInfoHandler)
	userGroup.POST("password", api.UserPasswordHandler)
	userGroup.POST("logout", api.UserLogoutHandler)
	userGroup.POST("captcha", api.UserCaptchaHandler)
	userGroup.GET("verify", api.UserVerifyHandler)

	normalGroup := apiGroup2.Group("normal")
	normalGroup.Use(middleware.JWTAuthMiddleware())
	normalGroup.GET("/problems/:id", api.NormalGetProblemByIdHandler)
	normalGroup.GET("/problems/slug/:slug", api.NormalGetProblemBySlugHandler)
	normalGroup.GET("/problems/list/:difficulty", api.NormalGetProblemListHandler)

	adminGroup := apiGroup2.Group("admin")
	adminGroup.Use(middleware.JWTAuthMiddleware())
	adminGroup.Use(middleware.JWTAdminAuthMiddleware())
	adminGroup.POST("/problems", api.AdminProblemSetHandler)
	adminGroup.PUT("/problems/:id", api.AdminProblemUpdateHandler)

	return r
}
