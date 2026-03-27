package model

type Attribute struct {
	IsWord bool `json:"isWord"`
	Word  string  `json:"word"`
	Score float64 `json:"score"`
}
