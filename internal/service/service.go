package service

import (
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository"
	"github.com/denis-shcherbinin/spbpu-software-design-project/pkg/hasher"
)

type Auth interface {
	SignUp(opts SignUpOpts) error
	SignIn(opts SignInOpts) (string, string, error)
}

type User interface {
	GetIDByCredentials(username, passwordHash string) (int64, error)
}

type List interface {
	Create(userID int64, opts CreateListOpts) error
	GetAll(userID int64) ([]domain.List, error)
	GetByID(userID, listID int64) (*domain.List, error)
	Update(userID, listID int64, opts UpdateListOpts) error
	DeleteByID(userID, listID int64) error
}

type Item interface {
	Create(userID, listID int64, opts CreateItemOpts) error
	GetAll(userID, listID int64) ([]domain.Item, error)
	GetByID(userID, itemID int64) (*domain.Item, error)
	Update(userID, itemID int64, opts UpdateItemOpts) error
	DeleteByID(userID, itemID int64) error
}

type Service struct {
	Auth Auth
	User User
	List List
	Item Item
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

	itemSvc := NewItemService(opts.Repo.Item, opts.Repo.List)

	return &Service{
		Auth: authSvc,
		User: NewUserService(opts.Repo.User),
		List: NewListService(opts.Repo.List),
		Item: itemSvc,
	}
}
