package entity

import "database/sql"

type List struct {
	ID          int64        `db:"id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	CreatedAt   sql.NullTime `db:"created_at"`
}
