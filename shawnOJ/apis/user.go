package apis

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"shawnOJ/dao"
	"shawnOJ/midware"
	"shawnOJ/model"
	"shawnOJ/utils"
	"time"
)

func Register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.ResponsesFail(c, "empty information")
		return
	}
	username := c.PostForm("username")
	password, err := midware.GenerateFormPassword(c.PostForm("password"))
	if err != nil {
		utils.ResponsesFail(c, "generate password wrong")
		return
	}
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	flag := dao.IfUserExist(username)
	if flag {
		utils.ResponsesFail(c, "user is exist")
		return
	}
	dao.AddUser(phone, email, username, password)
	utils.ResponsesSuccess(c, "register success")
}

func Login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.ResponsesFail(c, "empty information")
		return
	}
	username := c.PostForm("username")
	password, err := midware.GenerateFormPassword(c.PostForm("password"))
	if err != nil {
		utils.ResponsesFail(c, "generate password wrong")
		return
	}
	flag := dao.IfUserExist(username)
	if !flag {
		utils.ResponsesFail(c, "user not found")
		return
	}
	hashPassword := dao.CheckHashPassword(username)
	if midware.CompareHashAndPassword(hashPassword, password) {
		utils.ResponsesFail(c, "invalid password")
		return
	}
	claim := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "shawn",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(midware.Secret)
	utils.ResponsesSuccess(c, tokenString)
}

func SendEmail(c *gin.Context) {
	email := c.PostForm("email")

	// 先检查 Redis 里是否已有验证码
	_, err := utils.RedisClient.Get(utils.Ctx, "email_code:"+email).Result()
	if err == nil {
		utils.ResponsesFail(c, "You can only request a verification code once per minute")
		return
	}

	// 生成 6 位随机验证码
	code := midware.SendVerificationEmail(email)

	// 存到 Redis，设置过期时间为 5 分钟
	utils.RedisClient.Set(utils.Ctx, "email_code:"+email, code, 5*time.Minute)
	utils.RedisClient.Set(utils.Ctx, "email_limit:"+email, "1", 1*time.Minute) // 1 分钟限流

	utils.ResponsesSuccess(c, "verification code sent")
}

func VerifyEmail(c *gin.Context) {
	email := c.PostForm("email")
	codeClient := c.PostForm("code")

	if email == "" || codeClient == "" {
		utils.ResponsesFail(c, "email and code cannot be empty")
		return
	}

	// 从 Redis 获取验证码
	codeServer, err := utils.RedisClient.Get(utils.Ctx, "email_code:"+email).Result()
	if err == redis.Nil {
		utils.ResponsesFail(c, "verification code expired or not found")
		return
	} else if err != nil {
		utils.ResponsesFail(c, "failed to verify code")
		return
	}

	// 比较验证码
	if codeClient != codeServer {
		utils.ResponsesFail(c, "invalid verification code")
		return
	}

	// 验证成功后删除验证码，防止重复使用
	utils.RedisClient.Del(utils.Ctx, "email_code:"+email)

	utils.ResponsesSuccess(c, "email verification successful")
}
