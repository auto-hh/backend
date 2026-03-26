package service

import "github.com/auto-hh/backend/internal/repository"

type LLM struct {
	repository repository.IProfile
}

func NewLLM(repository repository.IProfile) ILLM {
	return &LLM{repository}
}

func (llm *LLM) FindVacancies() error {

	return nil
}
