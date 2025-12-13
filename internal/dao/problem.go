package dao

import (
	"ShawnOJ/internal"
	"ShawnOJ/internal/models"
	"ShawnOJ/utils/database"
	"database/sql"
	"errors"
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

func GetProblemNumber() (count int64, err error) {
	sqlStr := "select count(*) from `problems`"
	if err = database.Db.Get(&count, sqlStr); err != nil {
		return
	}
	return count, nil
}

func SelectProblemListWithPage(whereCondition string, args []interface{}, limit, offset int64) (list []models.Problem, total int64, err error) {
	// 1. 先查询总数（用于前端计算总页数）
	countSQL := `select count(*) from problems`
	if whereCondition != "" {
		if len(args) == 0 {
			return nil, 0, internal.ErrParamWrong
		}
		countSQL += " where " + whereCondition
	}
	// 执行总数查询（复用原有错误处理风格）
	err = database.Db.Get(&total, countSQL, args...)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, 0, err
		}
		total = 0 // 无数据时总数为0
	}

	// 2. 拼接分页查询SQL（完全复用原有列表查询的字段和排序）
	querySQL := `
		select id, title, difficulty, is_public, pass_rate, slug
		from problems 
	`
	if whereCondition != "" {
		querySQL += "where " + whereCondition + " "
	}
	// 保留原有排序 + 拼接分页条件（LIMIT/OFFSET）
	querySQL += `order by id asc limit ? offset ?`

	// 3. 拼接参数：原有WHERE参数 + limit + offset
	queryArgs := append(args, limit, offset)

	// 4. 执行分页查询（复用原有错误处理逻辑）
	var problems []models.Problem
	err = database.Db.Select(&problems, querySQL, queryArgs...)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, 0, err
		}
	}

	return problems, total, nil
}
