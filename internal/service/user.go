package service

import "github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"

type UserService struct {
	UserRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}
