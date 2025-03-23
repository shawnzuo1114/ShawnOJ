package apis

import (
	"github.com/gin-gonic/gin"
	"shawnOJ/dao"
	"shawnOJ/utils"
)

func UploadAvatar(c *gin.Context) {
	uname, _ := c.Get("username")
	username := uname.(string)
	avatar := c.PostForm("avatar")
	flag1 := dao.IfUserExist(username)
	if !flag1 {
		utils.ResponsesFail(c, "user is not exist")
		return
	}
	flag2 := dao.IfAvatarExist(username)
	if flag2 {
		dao.UpdateAvatar(username, avatar)
		utils.ResponsesSuccess(c, "update success")
	} else {
		dao.AddAvatar(username, avatar)
		utils.ResponsesSuccess(c, "upload success")
	}
}

func GetAvatar(c *gin.Context) {
	uname, _ := c.Get("username")
	username := uname.(string)
	var avatar string

	flag := dao.IfAvatarExist(username)
	if flag {
		avatar = dao.GetAvatar(username)
		avatarUrl := "http://localhost:8088/assets/" + avatar
		utils.ResponsesSuccess(c, avatarUrl)
	} else {
		utils.ResponsesFail(c, "user is not exist")
	}
}
