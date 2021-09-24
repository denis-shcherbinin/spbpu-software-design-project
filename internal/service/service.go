package service

import (
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
	"github.com/denis-shcherbinin/spbpu-software-design-project/pkg/hasher"
)

type Auth interface {
	SignUp(opts SignUpOpts) (*domain.User, error)
	SignIn(opts SignInOpts) (string, error)
}

type User interface {
}

type Service struct {
	Auth Auth
	User User
}

func NewService(repo *repository.Repository, hasher hasher.Hasher) *Service {
	authSvc := NewAuthService(NewAuthOpts{
		AuthRepo: repo.Auth,
		UserRepo: repo.User,
		Hasher:   hasher,
	})

	return &Service{
		Auth: authSvc,
		User: NewUserService(repo.User),
	}
}
