package dao

import (
	"ShawnOJ/internal"
	"ShawnOJ/internal/models"
	"ShawnOJ/utils/database"
)

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

func SelectUserPassword(user *models.User) (password string, err error) {
	sqlStr := "select username, userid, password from user where username = ?"
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
