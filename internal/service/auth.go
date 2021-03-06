package service

import (
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
	"github.com/denis-shcherbinin/spbpu-software-design-project/pkg/hasher"
)

type AuthService struct {
	AuthRepo repository.Auth
	UserRepo repository.User
	Hasher   hasher.Hasher
}

type NewAuthOpts struct {
	AuthRepo repository.Auth
	UserRepo repository.User
	Hasher   hasher.Hasher
}

func NewAuthService(opts NewAuthOpts) *AuthService {
	return &AuthService{
		AuthRepo: opts.AuthRepo,
		UserRepo: opts.UserRepo,
		Hasher:   opts.Hasher,
	}
}

type SignUpOpts struct {
	FirstName  string
	SecondName string
	Username   string
	Password   string
}

// SignUp .
func (svc *AuthService) SignUp(opts SignUpOpts) error {
	err := svc.AuthRepo.CreateUser(repository.CreateUserOpts{
		FirstName:  opts.FirstName,
		SecondName: opts.SecondName,
		Username:   opts.Username,
		Password:   svc.Hasher.Hash(opts.Password),
	})
	if err != nil {
		return err
	}

	return nil
}

type SignInOpts struct {
	Username string
	Password string
}

// SignIn .
func (svc *AuthService) SignIn(opts SignInOpts) (string, string, error) {
	passwordHash := svc.Hasher.Hash(opts.Password)

	ok, err := svc.UserRepo.CheckByCredentials(opts.Username, passwordHash)
	if err != nil {
		return "", "", err
	}
	if !ok {
		return "", "", errs.ErrUserNotFound
	}

	return opts.Username, passwordHash, nil
}
