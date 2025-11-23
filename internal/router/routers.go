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
	apiGroup2.Use(middleware.JWTAuthMiddleware())
	apiGroup2.GET("user/me", api.UserInfoHandler)
	apiGroup2.POST("user/password", api.UserPasswordHandler)
	apiGroup2.POST("user/logout", api.UserLogoutHandler)

	return r
}
