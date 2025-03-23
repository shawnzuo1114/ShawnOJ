package apis

import (
	"github.com/gin-gonic/gin"
	"shawnOJ/midware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(midware.CORS())
	r.Static("/assets", "./assets")
	r.POST("/shawnOJ", Login)
	r.PUT("/shawnOJ", Register)
	r.PUT("/shawnOJ/email", VerifyEmail)
	r.POST("/shawnOJ/email", SendEmail)
	UserRouter := r.Group("/shawnOJ")
	{
		UserRouter.Use(midware.JWTAuthMiddleware())
		UserRouter.POST("/avatar", UploadAvatar)
		UserRouter.GET("/avatar", GetAvatar)
	}
	r.Run(":8088")
}
