package dao

import (
	"ShawnOJ/internal/models"
	"ShawnOJ/utils/database"
)

func InsertProblem(p *models.Problem) (err error) {
	sqlStr := "insert into `problems` " +
		"(`title`, `slug`, `description`, `input_desc`, `output_desc`, `time_limit`, " +
		"`memory_limit`, `difficulty`, `creator_id`, `is_public`) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	if _, err = database.Db.Exec(sqlStr, p.Title, p.Slug, p.Description,
		p.InputDesc, p.OutputDesc, p.TimeLimit, p.MemoryLimit, p.Difficulty,
		p.CreatorID, p.IsPublic); err != nil {
		return
	}
	return nil
}

func SelectProblemById(p *models.Problem) (err error) {
	sqlStr := "select * from `problems` where `id` = ?"
	if err = database.Db.Get(p, sqlStr, p.ProblemID); err != nil {
		return
	}
	return nil
}

func SelectProblemBySlug(p *models.Problem) (err error) {
	sqlStr := "select * from `problems` where `slug` = ?"
	if err = database.Db.Get(p, sqlStr, p.Slug); err != nil {
		return
	}
	return nil
}
