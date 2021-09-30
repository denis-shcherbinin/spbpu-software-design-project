package domain

type User struct {
	ID           int64  `json:"id"`
	FullName     string `json:"full_name"`
	Username     string `json:"username"`
	RegisteredAt int64  `json:"registered_at"`
}
