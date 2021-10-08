package service

import (
	"fmt"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
)

type UserService struct {
	UserRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (svc *UserService) GetIDByCredentials(username, passwordHash string) (int64, error) {
	userID, err := svc.UserRepo.GetIDByCredentials(username, passwordHash)
	if err != nil {
		return 0, fmt.Errorf("UserService: %v", err)
	}

	return userID, nil
}
