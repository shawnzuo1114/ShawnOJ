package models

import "time"

type Problem struct {
	ProblemID   int64      `db:"id" json:"problem_id"`
	Title       string     `db:"title" json:"title"`
	Slug        string     `db:"slug" json:"slug"`
	Description string     `db:"description" json:"description"`
	InputDesc   string     `db:"input_desc" json:"input_desc"`
	OutputDesc  string     `db:"output_desc" json:"output_desc"`
	TimeLimit   int64      `db:"time_limit" json:"time_limit"`
	MemoryLimit int64      `db:"memory_limit" json:"memory_limit"`
	Difficulty  int8       `db:"difficulty" json:"difficulty"`
	CreatorID   int64      `db:"creator_id" json:"creator_id"`
	IsPublic    int8       `db:"is_public" json:"is_public"`
	SubmitCount int        `db:"submit_count" json:"submit_count"`
	AcceptCount int        `db:"accept_count" json:"accept_count"`
	PassRate    float64    `db:"pass_rate" json:"pass_rate"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}
