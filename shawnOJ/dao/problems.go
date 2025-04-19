package dao

import (
	"log"
	"shawnOJ/utils"
)

func GetProblem(id int) (string, string) {
	sqlStr := `SELECT title, description FROM problems WHERE id = ?`
	var title, description string
	err := utils.Db.QueryRow(sqlStr, id).Scan(&title, &description)
	if err != nil {
		log.Fatalln(err)
	}
	return title, description
}

func AddProblem() {

}
