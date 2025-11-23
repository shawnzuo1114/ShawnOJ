package middleware

import (
	"ShawnOJ/internal/api/responses"
	"ShawnOJ/utils/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userid"
const TokenKey = "token"

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			responses.ResponseError(c, responses.CodeNeedAuth)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			responses.ResponseError(c, responses.CodeInvalidAuth)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		token := parts[1]
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			responses.ResponseError(c, responses.CodeInvalidAuth)
			c.Abort()
			return
		}
		// 将当前请求的userid和token信息保存到请求的上下文c上
		c.Set(TokenKey, token)
		c.Set(CtxUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get("userid")来获取当前请求的用户信息
	}
}
