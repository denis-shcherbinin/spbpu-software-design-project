package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository/entity"
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

func (repo *AuthRepo) CreateUser(opts CreateUserOpts) (*domain.User, error) {
	const query = `
		INSERT INTO 
			t_user (first_name, second_name, username, password_hash)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			*`

	var user entity.User
	err := repo.DB.Get(&user, query,
		opts.FirstName,  // 1
		opts.SecondName, // 2
		opts.Username,   // 3
		opts.Password,   // 4
	)
	if err != nil {
		// TODO: error handling
		if err != sql.ErrNoRows {
			return nil, ErrUserAlreadyExists
		}

		return nil, err
	}

	return user.ToDomain(), nil
}
