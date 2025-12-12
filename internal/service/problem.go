package service

import (
	"ShawnOJ/internal"
	"ShawnOJ/internal/api/params"
	"ShawnOJ/internal/dao"
	"ShawnOJ/internal/models"
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
	return problemInfo, nil
}
