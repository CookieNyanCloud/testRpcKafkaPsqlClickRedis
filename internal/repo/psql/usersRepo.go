package psql

import (
	"context"
	"fmt"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type IUsersRepo interface {
	CreateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]domain.User, error)
}

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) IUsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) CreateUser(ctx context.Context, user domain.User) error {
	query := fmt.Sprintf("INSERT INTO users (id, name, password_hash) values ($1, $2)")
	_, err := r.db.Exec(query, user.Id, user.Name, user.PasswordHash)
	if err != nil {
		return errors.Wrap(err, "postgres creating user")
	}
	return nil
}

func (r *UsersRepo) DeleteUser(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id = $1")
	_, err := r.db.Exec(query, id)
	if err != nil {
		return errors.Wrap(err, "postgres deleting user")
	}
	return nil
}

func (r *UsersRepo) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	query := fmt.Sprintf("SELECT * FROM users")
	rows, err := r.db.Query(query)
	if err != nil {
		return []domain.User{}, errors.Wrap(err, "postgres query create")
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user)
		if err != nil {
			return []domain.User{}, errors.Wrap(err, "postgres single create")
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return []domain.User{}, errors.Wrap(err, "postgres rows create")

	}

	return users, nil
}
