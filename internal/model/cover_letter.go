package model

type GenerateRequest struct {
	Resume Profile `json:"resume"`
	Vacancy Vacancy `json:"vacancy"`
}

type CoverLetter struct {
	Letter string `json:"letter"`
	Status string `json:"status"`
}
