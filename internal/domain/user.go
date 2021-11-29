package domain

import "time"

type User struct {
	Id           string `db:"id" redis:"id"`
	Name         string `db:"name" redis:"name"`
	PasswordHash string `db:"password_hash" redis:"password_hash"`
}

type UserLog struct {
	Id        string    `db:"id" redis:"id"`
	Name      string    `db:"name" redis:"name"`
	CreatedAt time.Time `db:"created_at"`
}
