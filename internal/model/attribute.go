package model

type Attribute struct {
	IsWord bool    `json:"is_word"`
	Text   string  `json:"text"`
	Weight  float64 `json:"weight"`
}
