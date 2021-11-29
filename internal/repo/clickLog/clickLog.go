package clickLog

import (
	"fmt"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ClickLog struct {
	db *sqlx.DB
}

func NewClickHouseLog(db *sqlx.DB) IClickLog {
	return &ClickLog{db: db}
}

type IClickLog interface {
	LogNewUser(user domain.User) error
	ShowAllLogs() ([]domain.UserLog, error)
}

func (c *ClickLog) LogNewUser(user domain.User) error {
	query := fmt.Sprintf("INSERT INTO users (id, name) values ($1, $2)")
	_, err := c.db.Exec(query, user.Id, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClickLog) ShowAllLogs() ([]domain.UserLog, error) {
	var usersLog []domain.UserLog
	query := fmt.Sprintf("SELECT id, name, created_at FROM users")
	rows, err := c.db.Query(query)
	if err != nil {
		return []domain.UserLog{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var userLog domain.UserLog
		err := rows.Scan(userLog)
		if err != nil {
			return []domain.UserLog{}, err
		}
		usersLog = append(usersLog, userLog)
	}
	if err := rows.Err(); err != nil {
		return []domain.UserLog{}, err
	}
	return usersLog, nil
}
