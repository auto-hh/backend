package model

type Vacancy struct {
	JobTitle string `json:"job_title"`
	Salary   int `json:"salary"`
	City     string `json:"city"`
	WorkFormat string `json:"work_format"`
	Score     float64 `json:"score"`
	Link     string `json:"link"`
}
