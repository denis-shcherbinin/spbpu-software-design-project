package entity

import (
	"database/sql"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
)

type List struct {
	ID          int64        `db:"id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	CreatedAt   sql.NullTime `db:"created_at"`
}

func (l *List) ToDomain() *domain.List {
	var createdAt int64
	if l.CreatedAt.Valid {
		createdAt = l.CreatedAt.Time.Unix()
	}

	return &domain.List{
		ID:          l.ID,
		Title:       l.Title,
		Description: l.Description,
		CreatedAt:   createdAt,
	}
}
