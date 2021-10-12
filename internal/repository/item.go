package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/repository/entity"
)

type ItemRepo struct {
	DB *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) *ItemRepo {
	return &ItemRepo{
		DB: db,
	}
}

type CreateItemOpts struct {
	Title       string
	Description string
}

// Create creates item of list with passed id
// It returns internal errors.
func (repo *ItemRepo) Create(listID int64, opts CreateItemOpts) error {
	tx, err := repo.DB.Beginx()
	if err != nil {
		return fmt.Errorf("ItemRepo: %v", err)
	}

	itemQuery := `
		INSERT INTO
			t_item(title, description)
		VALUES
			($1, $2)
		RETURNING
			id`

	var itemID int64
	err = tx.QueryRow(itemQuery,
		opts.Title,
		opts.Description).
		Scan(&itemID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("ItemRepo: %v", err)
	}

	listItemQuery := `
		INSERT INTO
			t_list_item(list_id, item_id)
		VALUES
			($1, $2)`
	_, err = tx.Exec(listItemQuery, listID, itemID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("ItemRepo: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("ItemRepo: %v", err)
	}

	return nil
}

// GetAll form slice with all items of user list with passed id
// It returns all items or
// errs.ErrItemNotFound if listID or userID is wrong and internal errors.
func (repo *ItemRepo) GetAll(userID, listID int64) ([]entity.Item, error) {
	query := `
		SELECT
			i.id,
			i.title,
			i.description,
			i.done,
			i.created_at
		FROM
			t_item i
		INNER JOIN
			t_list_item li
		ON
			li.item_id = i.id
		INNER JOIN
			t_user_list ul
		ON
			ul.list_id = li.list_id
		WHERE
			li.list_id = $1
				AND
		    ul.user_id = $2
			`

	var items []entity.Item

	if err := repo.DB.Select(&items, query, listID, userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ItemRepo: %v", errs.ErrItemNotFound)
		}
		return nil, fmt.Errorf("ItemRepo: %v", err)
	}

	return items, nil
}

// GetByID returns user item with passed id or
// errs.ErrItemNotFound and internal errors.
func (repo *ItemRepo) GetByID(userID, itemID int64) (*entity.Item, error) {
	query := `
		SELECT
			i.id,
			i.title,
			i.description,
			i.done,
			i.created_at
		FROM
			t_item i
		INNER JOIN
			t_list_item li
		ON
			li.item_id = i.id
		INNER JOIN
			t_user_list ul
		ON
			ul.list_id = li.list_id
		WHERE
			i.id = $1
				AND
			ul.user_id = $2`

	var item entity.Item

	if err := repo.DB.Get(&item, query, itemID, userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ItemRepo: %v", errs.ErrItemNotFound)
		}
		return nil, fmt.Errorf("ItemRepo: %v", err)
	}

	return &item, nil
}

type UpdateItemOpts struct {
	Title       *string
	Description *string
	Done        *bool
}

// Update updates user item with passed id
// It returns errs.ErrItemNotFound if wrong item id and internal errors.
func (repo *ItemRepo) Update(userID, itemID int64, opts UpdateItemOpts) error {
	query := `
		UPDATE
			t_item i
		SET
			title 		= COALESCE($1, i.title),
		    description = COALESCE($2, i.description),
		    done 		= COALESCE($3, i.done)
		FROM
			t_list_item li,
		    t_user_list ul
		WHERE
			i.id = li.item_id
				AND
		    li.list_id = ul.list_id
				AND
		    ul.user_id = $4
				AND
		    i.id = $5`

	result, err := repo.DB.Exec(query,
		opts.Title,       // 1
		opts.Description, // 2
		opts.Done,        // 3
		userID,           // 4
		itemID,           // 5
	)

	if err != nil {
		return fmt.Errorf("ItemRepo: %v", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ItemRepo: %v", err)
	}
	if count != 1 {
		return fmt.Errorf("ItemRepo: %v", errs.ErrItemNotFound)
	}

	return nil
}

func (repo *ItemRepo) DeleteByID(userID, itemID int64) error {
	query := `
		DELETE
		FROM
			t_item i
		USING
			t_list_item li,
			t_user_list ul
		WHERE
			li.item_id = i.id
				AND
			ul.list_id = li.list_id
				AND
			ul.user_id = $1
				AND
			i.id = $2`

	result, err := repo.DB.Exec(query, userID, itemID)
	if err != nil {
		return fmt.Errorf("ItemRepo: %v", err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ItemRepo: %v", err)
	}
	if count != 1 {
		return fmt.Errorf("ItemRepo: %v", errs.ErrItemNotFound)
	}

	return nil
}
