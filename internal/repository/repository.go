package repository

import (
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(opts CreateUserOpts) error
}

type User interface {
	GetIDByCredentials(username string, passwordHash string) (int64, error)
}

type Repository struct {
	Auth Auth
	User User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
		User: NewUserRepo(db),
	}
}
