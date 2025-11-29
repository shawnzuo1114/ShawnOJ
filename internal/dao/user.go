package dao

import (
	"ShawnOJ/internal"
	"ShawnOJ/internal/models"
	"ShawnOJ/utils/database"
	"time"

	"github.com/gin-gonic/gin"
)

const duration = 5 * time.Minute

func CheckUser(username string) (err error) {
	sqlStr := "select count(username) from user where username = ?"
	var count int
	if err = database.Db.Get(&count, sqlStr, username); err != nil {
		return
	}
	if count > 0 {
		return internal.ErrUserExists
	}
	return nil
}

func QueryUser(username string) (err error) {
	sqlStr := "select count(username) from user where username = ?"
	var count int
	if err = database.Db.Get(&count, sqlStr, username); err != nil {
		return
	}
	if count <= 0 {
		return internal.ErrUserNotExists
	}
	return nil
}

func InsertUser(u *models.User) (err error) {
	sqlStr := "insert into `user` (`userid`, `username`, `password`) values (?, ?, ?)"
	if _, err = database.Db.Exec(sqlStr, u.UserID, u.Username, u.Password); err != nil {
		return
	}
	return nil
}

func SelectUserInfo(user *models.User) (password string, err error) {
	sqlStr := "select username, userid, password, role from user where username = ?"
	if err = database.Db.Get(user, sqlStr, user.Username); err != nil {
		return
	}
	return user.Password, nil
}

func SelectUserInfoByID(user *models.User) (err error) {
	sqlStr := "select * from user where userid = ?"
	if err = database.Db.Get(user, sqlStr, user.UserID); err != nil {
		return
	}
	return
}

func CheckUserPassword(userid int64) (password string, err error) {
	sqlStr := "select password from user where userid = ?"
	if err = database.Db.Get(&password, sqlStr, userid); err != nil {
		return
	}
	return password, nil
}

func UpdatePassword(user *models.User) (err error) {
	sqlStr := "update `user` set `password` = ? where userid = ?"
	if _, err = database.Db.Exec(sqlStr, user.Password, user.UserID); err != nil {
		return
	}
	return nil
}

func GetCodeTTL(email string, ctx *gin.Context) (flag bool, err error) {
	c := ctx.Request.Context()
	key := "email:verify:code:" + email
	ttl, err := database.Rdb.TTL(c, key).Result()
	if err != nil {
		return false, err
	}
	return ttl > 0, nil
}

func SetVerifyCode(email string, code int, ctx *gin.Context) (err error) {
	c := ctx.Request.Context()
	key := "email:verify:code:" + email
	return database.Rdb.Set(c, key, code, duration).Err()
}

func GetVerifyCode(email string, ctx *gin.Context) (value string, err error) {
	c := ctx.Request.Context()
	key := "email:verify:code:" + email
	return database.Rdb.Get(c, key).Result()
}

func DeleteVerifyCode(email string, ctx *gin.Context) (err error) {
	c := ctx.Request.Context()
	key := "email:verify:code:" + email
	return database.Rdb.Del(c, key).Err()
}

func CreateEmail(email string, userid int64) (err error) {
	sqlStr := "UPDATE `user` SET email = ? WHERE userid = ?;"
	if _, err = database.Db.Exec(sqlStr, email, userid); err != nil {
		return
	}
	return nil
}

func AddTokenBlacklist(token string, expireTime time.Duration, ctx *gin.Context) (err error) {
	c := ctx.Request.Context()
	key := "blacklist:token:" + token
	return database.Rdb.Set(c, key, "1", duration).Err()
}

func IsTokenInBlacklist(token string, c *gin.Context) (bool, error) {
	exists, err := database.Rdb.Exists(c, "blacklist:token:"+token).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}
