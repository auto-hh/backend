package model

type Vacancy struct {
	JobTitle string `json:"jobTitle"`
	Salary   int `json:"salary"`
	City     string `json:"city"`
	WorkFormat string `json:"workFormat"`
	Score     float64 `json:"score"`
	Link     string `json:"link"`
}
