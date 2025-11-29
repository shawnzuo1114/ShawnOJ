package jwt

import (
	"ShawnOJ/internal"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("shawnzuo")

type MyClaims struct {
	UserID   int64  `json:"userid"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenToken(userID int64, username string, role string) (string, error) {
	c := MyClaims{
		userID,
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "shawnoj",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func GetUserID(c *gin.Context) (int64, error) {
	val, ok := c.Get("userid")
	if !ok {
		return 0, errors.New("userid not found")
	}

	switch v := val.(type) {
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	default:
		return 0, fmt.Errorf("invalid userid type: %T", val)
	}
}

func GetRemainingExpireTime(token string) (time.Duration, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return 0, err
	}

	expireTimestamp := claims.ExpiresAt
	expireTime := time.Unix(expireTimestamp, 0)
	now := time.Now()
	remaining := expireTime.Sub(now)
	if remaining < 0 {
		return 0, internal.ErrTokenExpired
	}
	return remaining, nil
}
