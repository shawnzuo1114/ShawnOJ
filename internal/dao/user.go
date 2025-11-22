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

func InsertUser(u *models.User) (err error) {
	sqlStr := "insert into `user` (`userid`, `username`, `password`) values (?, ?, ?)"
	if _, err = database.Db.Exec(sqlStr, u.UserID, u.Username, u.Password); err != nil {
		return
	}
	return nil
}
