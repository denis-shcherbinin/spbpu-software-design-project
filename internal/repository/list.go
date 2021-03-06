package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository/entity"
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
	Title       string
	Description string
}

// Create .
func (repo *ListRepo) Create(userID int64, opts CreateListOpts) error {
	tx, err := repo.DB.Beginx()
	if err != nil {
		return err
	}

	listQuery := `
		INSERT INTO
			t_list(title, description)
		VALUES
		    ($1, $2)
		RETURNING
			id`

	var listID int64
	err = tx.Get(&listID, listQuery, opts.Title, opts.Description)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	userListQuery := `
		INSERT INTO
			t_user_list(user_id, list_id)
		VALUES
		    ($1, $2)`

	_, err = tx.Exec(userListQuery, userID, listID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetAll .
func (repo *ListRepo) GetAll(userID int64) ([]entity.List, error) {
	query := `
		SELECT
			l.id, 
			l.title,
			l.description, 
			l.created_at
		FROM
			t_list l
		INNER JOIN
			t_user_list ul
		ON 
			l.id = ul.list_id
		WHERE
			ul.user_id = $1`

	var lists []entity.List

	err := repo.DB.Select(&lists, query, userID)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

// GetByID .
func (repo *ListRepo) GetByID(userID, listID int64) (*entity.List, error) {
	query := `
		SELECT
			l.id,
			l.title,
			l.description,
			l.created_at
		FROM
			t_list l
				INNER JOIN
			t_user_list ul
				ON
					l.id = ul.list_id
		WHERE
			ul.user_id = $1
				AND
			ul.list_id = $2`

	var list entity.List
	err := repo.DB.Get(&list, query, userID, listID)
	if err != nil {
		// list with such id doesn't exist
		if err == sql.ErrNoRows {
			return nil, errs.ErrListNotFound
		}
		return nil, err
	}

	return &list, nil
}

type UpdateListOpts struct {
	Title       *string
	Description *string
}

// Update .
func (repo *ListRepo) Update(userID, listID int64, opts UpdateListOpts) error {
	query := `
		UPDATE
			t_list l
		SET
			title 		  = COALESCE($1, title),
<<<<<<< HEAD
			description = COALESCE($2, description)
=======
			description   = COALESCE($2, description)
>>>>>>> 5343e6538e61773e7dba7298d408d5d379b7d030
		FROM 
			t_user_list ul
		WHERE
			l.id = ul.list_id
				AND
			ul.user_id = $3
				AND
			ul.list_id = $4`

	log.Debug("query:", query)
	log.Debug("set parameters:", opts.Title, opts.Description)

	result, err := repo.DB.Exec(query,
		opts.Title,       // 1
		opts.Description, // 2
		userID,           // 3
		listID,           // 4
	)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return errs.ErrListNotFound
	}

	return nil
}

// DeleteByID .
func (repo *ListRepo) DeleteByID(userID, listID int64) error {
	query := `
		DELETE FROM
			t_list l
		USING
			t_user_list ul
		WHERE
			l.id = ul.list_id
				AND
			ul.user_id = $1
				AND
			ul.list_id = $2`

	result, err := repo.DB.Exec(query, userID, listID)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return errs.ErrListNotFound
	}

	return nil
}
