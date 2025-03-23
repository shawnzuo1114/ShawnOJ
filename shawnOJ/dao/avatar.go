package dao

import (
	"log"
	"shawnOJ/utils"
)

func AddAvatar(username, avatar string) {
	sqlStr := "insert into avatar(username, avatar_url) values (?, ?)"
	_, err := utils.Db.Exec(sqlStr, username, avatar)
	if err != nil {
		log.Fatalln(err)
	}
}

func UpdateAvatar(username, avatar string) {
	sqlStr := "update avatar set avatar_url=? where username=?"
	_, err := utils.Db.Exec(sqlStr, avatar, username)
	if err != nil {
		log.Fatalln(err)
	}
}

func IfAvatarExist(username string) bool {
	sqlStr := "select count(*) from avatar where username=?"
	var count int
	err := utils.Db.QueryRow(sqlStr, username).Scan(&count)
	if err != nil {
		log.Fatalln(err)
	}
	return count > 0
}

func GetAvatar(username string) string {
	sqlStr := "select avatar_url from avatar where username=?"
	var avatarUrl string
	err := utils.Db.QueryRow(sqlStr, username).Scan(&avatarUrl)
	if err != nil {
		log.Fatalln(err)
	}
	return avatarUrl
}
