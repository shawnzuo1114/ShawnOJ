package service

import (
	"ShawnOJ/internal"
	"ShawnOJ/internal/api/params"
	"ShawnOJ/internal/dao"
	"ShawnOJ/internal/models"
	"strings"
)

func AdminProblemSetService(aps params.AdminProblemSetParam) (err error) {
	problem := &models.Problem{
		Title:       aps.Title,
		Slug:        aps.Slug,
		Description: aps.Description,
		InputDesc:   aps.InputDesc,
		OutputDesc:  aps.OutputDesc,
		TimeLimit:   aps.TimeLimit,
		MemoryLimit: aps.MemoryLimit,
		Difficulty:  aps.Difficulty,
		CreatorID:   aps.CreatorID,
		IsPublic:    aps.IsPublic,
	}

	if err = dao.InsertProblem(problem); err != nil {
		return
	}

	return nil
}

func NormalGetProblemByIdService(id int64, role string) (np params.NormalGetProblemParam, err error) {
	var problemInfo params.NormalGetProblemParam
	p := &models.Problem{
		ProblemID: id,
	}
	if err = dao.SelectProblemById(p); err != nil {
		return
	}

	if p.IsPublic == 0 && role == "user" {
		return params.NormalGetProblemParam{}, internal.ErrPermissionDenied
	}

	problemInfo.Title = p.Title
	problemInfo.Description = p.Description
	problemInfo.InputDesc = p.InputDesc
	problemInfo.OutputDesc = p.OutputDesc
	problemInfo.TimeLimit = p.TimeLimit
	problemInfo.MemoryLimit = p.MemoryLimit
	problemInfo.Difficulty = p.Difficulty
	problemInfo.SubmitCount = p.SubmitCount
	problemInfo.AcceptCount = p.AcceptCount
	problemInfo.PassRate = p.PassRate
	return problemInfo, nil
}

func NormalGetProblemBySlugService(slug string, role string) (np params.NormalGetProblemParam, err error) {
	var problemInfo params.NormalGetProblemParam
	p := &models.Problem{
		Slug: slug,
	}
	if err = dao.SelectProblemBySlug(p); err != nil {
		return
	}

	if p.IsPublic == 0 && role == "user" {
		return params.NormalGetProblemParam{}, internal.ErrPermissionDenied
	}

	problemInfo.Title = p.Title
	problemInfo.Description = p.Description
	problemInfo.InputDesc = p.InputDesc
	problemInfo.OutputDesc = p.OutputDesc
	problemInfo.TimeLimit = p.TimeLimit
	problemInfo.MemoryLimit = p.MemoryLimit
	problemInfo.Difficulty = p.Difficulty
	problemInfo.SubmitCount = p.SubmitCount
	problemInfo.AcceptCount = p.AcceptCount
	problemInfo.PassRate = p.PassRate
	return problemInfo, nil
}

func NormalGetProblemListService(role string, page int64, size int64, difficulty int64) (pl []params.NormalGetProblemListParam, err error) {
	if page < 1 || size < 1 {
		return nil, internal.ErrParamWrong
	}

	if size > 50 {
		return nil, internal.ErrParamWrong
	}
	offset := (page - 1) * size

	pl = make([]params.NormalGetProblemListParam, 0)
	var whereCondition []string
	var args []interface{}
	if role == "user" {
		whereCondition = append(whereCondition, "is_public = ?")
		args = append(args, 1)
	}

	if difficulty > 0 {
		whereCondition = append(whereCondition, "difficulty = ?")
		args = append(args, difficulty)
	}

	var whereStr string
	if len(whereCondition) > 0 {
		whereStr = strings.Join(whereCondition, " AND ")
	}

	problemList, total, err := dao.SelectProblemListWithPage(whereStr, args, size, offset)
	if err != nil {
		return nil, err
	}
	if total > 0 && len(problemList) == 0 {
		return nil, internal.ErrParamWrong
	}

	for _, p := range problemList {
		problemInfo := &params.NormalGetProblemListParam{
			Id:         p.ProblemID,
			Title:      p.Title,
			Difficulty: p.Difficulty,
			IsPublic:   p.IsPublic,
			PassRate:   p.PassRate,
		}
		pl = append(pl, *problemInfo)
	}
	return pl, nil
}
