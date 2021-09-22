package repository

import "github.com/jmoiron/sqlx"

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}
