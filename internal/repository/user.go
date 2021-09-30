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
	const query = `
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
	quotedUsername := fmt.Sprintf("'%s'", username)
	quotedPasswordHash := fmt.Sprintf("'%s'", passwordHash)
	err := repo.DB.Get(&exists, query, quotedUsername, quotedPasswordHash)
	if err != nil {
		return false, err
	}

	return exists, nil
}
