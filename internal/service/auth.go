package service

import (
	"encoding/base64"
	"fmt"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
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

type SignInOpts struct {
	Username string
	Password string
}

func (svc *AuthService) SignIn(opts SignInOpts) (string, error) {
	passwordHash := svc.Hasher.Hash(opts.Password)

	userID, err := svc.UserRepo.GetIDByCredentials(opts.Username, passwordHash)
	if err != nil {
		return "", err
	}

	token, err := svc.CreateSession(userID, passwordHash)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func (svc *AuthService) CreateSession(userID int64, passwordHash string) (string, error) {
	token := svc.createToken(userID, passwordHash)

	err := svc.AuthRepo.CreateSession(userID, token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (svc *AuthService) createToken(userID int64, passwordHash string) string {
	stringToEncode := fmt.Sprintf("%d:%s", userID, passwordHash)

	return base64.StdEncoding.EncodeToString([]byte(stringToEncode))
}

