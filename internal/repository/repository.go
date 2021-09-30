package repository

import (
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(opts CreateUserOpts) error
}

type User interface {
	CheckByCredentials(username, passwordHash string) (bool, error)
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
