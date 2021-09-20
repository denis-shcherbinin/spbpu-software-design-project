package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Opts struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

func NewPostgresDB(opts Opts) (*sqlx.DB, error) {
	addr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		opts.User,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.Name,
	)

	db, err := sqlx.Connect("postgres", addr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
