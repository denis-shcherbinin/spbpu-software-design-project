package repository

import (
	"github.com/jmoiron/sqlx"
)

type ListRepo struct {
	DB *sqlx.DB
}

func NewListRepo(db *sqlx.DB) *ListRepo {
	return &ListRepo{
		DB: db,
	}
}

type CreateListOpts struct {
	UserID      int64
	Title       string
	Description string
}

func (repo *ListRepo) Create(opts CreateListOpts) error {
	tx, err := repo.DB.Beginx()
	if err != nil {
		return err
	}

	const listQuery = `
		INSERT INTO
			t_list(title, description)
		VALUES($1, $2)
		RETURNING
			id`

	var listID int64
	err = tx.Get(&listID, listQuery, opts.Title, opts.Description)
	if err != nil {
		tx.Rollback()
		return err
	}

	const userListQuery = `
		INSERT INTO
			t_user_list(user_id, list_id)
		VALUES($1, $2)`

	_, err = tx.Exec(userListQuery, opts.UserID, listID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
