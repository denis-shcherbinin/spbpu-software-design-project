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

func (repo *UserRepo) GetIDByCredentials(username string, passwordHash string) (int64, error) {
	const query = `
		SELECT
			id
		FROM
			t_user
		WHERE
			username=$1
				AND
			password_hash=$2`

	var userID int64

	row := repo.DB.QueryRow(query, username, passwordHash)
	err := row.Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
