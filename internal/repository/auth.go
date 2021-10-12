package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
)

type AuthRepo struct {
	DB *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		DB: db,
	}
}

type CreateUserOpts struct {
	FirstName  string
	SecondName string
	Username   string
	Password   string
}

// CreateUser .
func (repo *AuthRepo) CreateUser(opts CreateUserOpts) error {
	query := `
		INSERT INTO 
			t_user (first_name, second_name, username, password_hash)
		VALUES
			($1, $2, $3, $4)`

	_, err := repo.DB.Exec(query,
		opts.FirstName,  // 1
		opts.SecondName, // 2
		opts.Username,   // 3
		opts.Password,   // 4
	)
	if err != nil {
		// User with passed username already exists
		if err != sql.ErrNoRows {
			return errs.ErrUserAlreadyExists
		}

		return err
	}

	return nil
}
