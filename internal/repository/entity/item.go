package entity

import (
	"database/sql"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
)

type Item struct {
	ID          int64        `db:"id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	Done        bool         `db:"done"`
	CreatedAt   sql.NullTime `db:"created_at"`
}

func (i *Item) ToDomain() *domain.Item {
	var createdAt int64
	if i.CreatedAt.Valid {
		createdAt = i.CreatedAt.Time.Unix()
	}

	return &domain.Item{
		ID:          i.ID,
		Title:       i.Title,
		Description: i.Description,
		Done:        i.Done,
		CreatedAt:   createdAt,
	}
}
