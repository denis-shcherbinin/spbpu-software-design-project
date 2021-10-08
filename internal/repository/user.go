package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) CheckByCredentials(username, passwordHash string) (bool, error) {
	query := `
		SELECT
			EXISTS
			(
				SELECT 
					COUNT(id)
				FROM
					t_user
				WHERE
					username = $1
						AND
					password_hash = $2
			) AS exists`

	var exists bool

	err := repo.DB.QueryRow(query,
		username,
		passwordHash).
		Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("UserRepo: %v", err)
	}

	return exists, nil
}

func (repo *UserRepo) GetIDByCredentials(username, passwordHash string) (int64, error) {
	query := `
			SELECT 
				id
			FROM
				t_user
			WHERE
				username = $1
					AND
				password_hash = $2`

	var id int64

	err := repo.DB.QueryRow(query,
		username,
		passwordHash).
		Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("UserRepo: %v", err)
	}

	return id, nil
}
