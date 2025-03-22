package apis

import (
	"github.com/gin-gonic/gin"
	"shawnOJ/midware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(midware.CORS())
	r.POST("/shawnOJ", Login)
	r.PUT("/shawnOJ", Register)
	r.PUT("/shawnOJ/email", VerifyEmail)
	r.GET("/shawnOJ/email", SendEmail)
	r.Run(":8080")
}
