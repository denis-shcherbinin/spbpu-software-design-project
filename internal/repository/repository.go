package repository

import (
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(opts CreateUserOpts) error
}

type User interface {
	CheckByCredentials(username, passwordHash string) (bool, error)
	GetIDByCredentials(username, passwordHash string) (int64, error)
}

type List interface {
	Create(opts CreateListOpts) error
}

type Repository struct {
	Auth Auth
	User User
	List List
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
		User: NewUserRepo(db),
		List: NewListRepo(db),
	}
}
