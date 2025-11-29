package middleware

import (
	"ShawnOJ/internal/api/responses"
	"ShawnOJ/internal/dao"
	"ShawnOJ/utils/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userid"
const TokenKey = "token"
const RoleKey = "role"

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

		inBlacklist, err := dao.IsTokenInBlacklist(token, c)
		if err != nil {
			responses.ResponseError(c, responses.CodeServerBusy)
			c.Abort()
			return
		}
		if inBlacklist {
			responses.ResponseError(c, responses.CodeInvalidAuth)
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			responses.ResponseError(c, responses.CodeInvalidAuth)
			c.Abort()
			return
		}
		// 将当前请求的userid和token信息保存到请求的上下文c上
		c.Set(TokenKey, token)
		c.Set(CtxUserIDKey, mc.UserID)
		c.Set(RoleKey, mc.Role)
		c.Next() // 后续的处理函数可以用过c.Get("userid")来获取当前请求的用户信息
	}
}

func JWTAdminAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		role, ok := c.Get(RoleKey)
		if !ok {
			responses.ResponseError(c, responses.CodeNeedAuth)
			c.Abort()
			return
		}

		if role.(string) != "admin" {
			responses.ResponseError(c, responses.CodeNotAdmin)
			c.Abort()
			return
		}

		c.Next()
	}
}
