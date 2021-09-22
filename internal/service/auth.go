package service

import (
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
	"github.com/denis-shcherbinin/spbpu-software-design-project/pkg/hasher"
)

type AuthService struct {
	AuthRepo repository.Auth
	Hasher   hasher.Hasher
}

func NewAuthService(authRepo repository.Auth, hasher hasher.Hasher) *AuthService {
	return &AuthService{
		AuthRepo: authRepo,
		Hasher:   hasher,
	}
}

type SignUpOpts struct {
	FirstName  string
	SecondName string
	Username   string
	Password   string
}

func (svc *AuthService) SignUp(opts SignUpOpts) (*domain.User, error) {
	user, err := svc.AuthRepo.CreateUser(repository.CreateUserOpts{
		FirstName:  opts.FirstName,
		SecondName: opts.SecondName,
		Username:   opts.Username,
		Password:   svc.Hasher.Hash(opts.Password),
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
