package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email"`
	Phone    string `form:"phone" json:"phone"`
	Status   string `form:"status" json:"status"`
	Created  string `form:"created" json:"created"`
	Updated  string `form:"updated" json:"updated"`
}
