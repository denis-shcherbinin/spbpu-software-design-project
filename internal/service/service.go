package service

import (
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
	"github.com/denis-shcherbinin/spbpu-software-design-project/pkg/hasher"
)

type Auth interface {
	SignUp(opts SignUpOpts) (*domain.User, error)
}

type User interface {
}

type Service struct {
	Auth Auth
	User User
}

func NewService(repo *repository.Repository, hasher hasher.Hasher) *Service {
	return &Service{
		Auth: NewAuthService(repo.Auth, hasher),
		User: NewUserService(repo.User),
	}
}
