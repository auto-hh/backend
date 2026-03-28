package model

type Vacancy struct {
	JobTitle   string  `json:"job_title"`
	Salary     string     `json:"salary"`
	City       string  `json:"city"`
	Body string `json:"body"`
	WorkFormat string  `json:"work_format"`
	Score      float64 `json:"score"`
	Link       string  `json:"link"`
}
