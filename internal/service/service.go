package service

import "github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
