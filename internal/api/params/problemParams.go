package params

type AdminProblemSetParam struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	InputDesc   string `json:"input_desc"`
	OutputDesc  string `json:"output_desc"`
	TimeLimit   int64  `json:"time_limit"`
	MemoryLimit int64  `json:"memory_limit"`
	Difficulty  int8   `json:"difficulty"`
	CreatorID   int64  `json:"creator_id"`
	IsPublic    int8   `json:"is_public"`
}

type NormalGetProblemParam struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	InputDesc   string  `json:"input_desc"`
	OutputDesc  string  `json:"output_desc"`
	TimeLimit   int64   `json:"time_limit"`
	MemoryLimit int64   `json:"memory_limit"`
	Difficulty  int8    `json:"difficulty"`
	SubmitCount int     `json:"submit_count"`
	AcceptCount int     `json:"accept_count"`
	PassRate    float64 `json:"pass_rate"`
}

type NormalGetProblemListParam struct {
	Id         int64   `json:"id"`
	Title      string  `json:"title"`
	Difficulty int8    `json:"difficulty"`
	IsPublic   int8    `json:"is_public"`
	PassRate   float64 `json:"pass_rate"`
}
