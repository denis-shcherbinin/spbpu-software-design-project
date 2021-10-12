package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository/entity"
)

type Auth interface {
	CreateUser(opts CreateUserOpts) error
}

type User interface {
	CheckByCredentials(username, passwordHash string) (bool, error)
	GetIDByCredentials(username, passwordHash string) (int64, error)
}

type List interface {
	Create(userID int64, opts CreateListOpts) error
	GetAll(userID int64) ([]entity.List, error)
	GetByID(userID, listID int64) (*entity.List, error)
	Update(userID, listID int64, opts UpdateListOpts) error
	DeleteByID(userID, listID int64) error
}

type Item interface {
	Create(listID int64, opts CreateItemOpts) error
	GetAll(userID int64, listID int64) ([]entity.Item, error)
	GetByID(userID, itemID int64) (*entity.Item, error)
	Update(userID, itemID int64, opts UpdateItemOpts) error
	DeleteByID(userID, listID int64) error
}

type Repository struct {
	Auth Auth
	User User
	List List
	Item Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
		User: NewUserRepo(db),
		List: NewListRepo(db),
		Item: NewItemRepo(db),
	}
}
