package service

import (
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
	"github.com/denis-shcherbinin/spbpu-software-design-project/pkg/hasher"
)

type Auth interface {
	SignUp(opts SignUpOpts) error
	SignIn(opts SignInOpts) (string, string, error)
	CheckByCredentials(username, passwordHash string) (bool, error)
}

type User interface {
}

type Service struct {
	Auth Auth
	User User
}

type NewServiceOpts struct {
	Repo   *repository.Repository
	Hasher hasher.Hasher
}

func NewService(opts NewServiceOpts) *Service {
	authSvc := NewAuthService(NewAuthOpts{
		AuthRepo: opts.Repo.Auth,
		UserRepo: opts.Repo.User,
		Hasher:   opts.Hasher,
	})

	return &Service{
		Auth: authSvc,
		User: NewUserService(opts.Repo.User),
	}
}
