package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
)

type Auth interface {
	CreateUser(opts CreateUserOpts) (*domain.User, error)
}

type User interface {
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
