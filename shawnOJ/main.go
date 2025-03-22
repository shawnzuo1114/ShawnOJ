package main

import (
	"shawnOJ/apis"
	"shawnOJ/utils"
)

func main() {
	utils.InitDB()
	utils.InitRedis()
	apis.InitRouter()
}
