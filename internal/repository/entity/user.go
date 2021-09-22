package entity

import (
	"database/sql"
	"strings"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/domain"
)

type User struct {
	ID           int64        `db:"id"`
	FirstName    string       `db:"first_name"`
	SecondName   string       `db:"second_name"`
	Username     string       `db:"username"`
	PasswordHash string       `db:"password_hash"`
	RegisteredAt sql.NullTime `db:"registered_at"`
}

func (u *User) ToDomain() *domain.User {
	fullName := strings.Join([]string{u.FirstName, u.SecondName}, " ")

	var registeredAt int64
	if u.RegisteredAt.Valid {
		registeredAt = u.RegisteredAt.Time.Unix()
	}

	return &domain.User{
		ID:           u.ID,
		FullName:     fullName,
		Username:     u.Username,
		RegisteredAt: registeredAt,
	}
}
