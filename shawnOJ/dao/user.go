package dao

import (
	"log"
	"shawnOJ/midware"
	"shawnOJ/utils"
)

func AddUser(phone, email, username, password string) {
	sqlStr := "insert into users(username, password_hash, phone, email) values (?, ?, ?, ?)"
	_, err := utils.Db.Exec(sqlStr, username, password, phone, email)
	if err != nil {
		log.Fatalln(err)
	}
}

func DeleteUser(info string) {
	sqlStr := "delete from users where username = ? or phone = ? or email = ?"
	_, err := utils.Db.Exec(sqlStr, info, info, info)
	if err != nil {
		log.Fatalln(err)
	}
}

func IfUserExist(info string) bool {
	sqlStr := "SELECT username FROM users WHERE username = ?"
	var s string
	err := utils.Db.QueryRow(sqlStr, info).Scan(&s)

	if err != nil {
		return false
	}

	return true
}

func UpdateUser(info string, password string) {
	sqlStr := "update users set password_hash = ? where username = ? or phone = ? or email = ?"
	passwordHash, err := midware.GenerateFormPassword(password)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = utils.Db.Exec(sqlStr, passwordHash, info, info)
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckHashPassword(username string) string {
	sqlStr := "select password_hash from users where username = ?"
	var passwordHash string
	err := utils.Db.QueryRow(sqlStr, username).Scan(&passwordHash)
	if err != nil {
		log.Fatalln(err)
	}
	return passwordHash
}

func GetUserInfo(username string) (e string, p string, s string) {
	sqlStr := "select email, phone, status from users where username = ?"
	var email string
	var phone string
	var status string
	err := utils.Db.QueryRow(sqlStr, username).Scan(&email, &phone, &status)
	if err != nil {
		log.Fatalln(err)
	}
	return email, phone, status
}
