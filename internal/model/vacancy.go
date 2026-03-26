package model

type Vacancy struct {
	JobTitle string `json:"jobTitle"`
	Salary   string `json:"salary"`
	City     string `json:"city"`
	Body     string `json:"body"`
	Link     string `json:"link"`
}
