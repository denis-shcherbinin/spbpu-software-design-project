package repository

import (
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
	row := repo.DB.QueryRow(query, username, passwordHash)
	if err := row.Scan(&exists); err != nil {
		return false, err
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
	row := repo.DB.QueryRow(query, username, passwordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
