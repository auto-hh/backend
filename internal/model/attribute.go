package model

type Attribute struct {
	IsWord bool    `json:"is_word"`
	Word   string  `json:"word"`
	Score  float64 `json:"score"`
}
