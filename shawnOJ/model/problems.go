package model

type Problems struct {
	Id                int    `json:"id" form:"id"`
	Title             string `json:"title" form:"title"`
	Description       string `json:"description" form:"description"`
	InputDescription  string `json:"input_description" form:"input_description"`
	OutputDescription string `json:"output_description" form:"output_description"`
	SampleInput       string `json:"sample_input" form:"sample_input"`
	SampleOutput      string `json:"sample_output" form:"sample_output"`
	TimeLimit         int    `json:"time_limit" form:"time_limit"`
	MemoryLimit       int    `json:"memory_limit" form:"memory_limit"`
	Difficulty        int    `json:"difficulty" form:"difficulty"`
}
