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

	apiGroup := r.Group("/api/v1")

	apiGroup.POST("user/register", api.UserRegisterHandler)

	return r
}
