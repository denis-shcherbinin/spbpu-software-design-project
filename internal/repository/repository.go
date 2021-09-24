package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
)

type Auth interface {
	CreateUser(opts CreateUserOpts) (*domain.User, error)
	CreateSession(userID int64, token string) error
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
