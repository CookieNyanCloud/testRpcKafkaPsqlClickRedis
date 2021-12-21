package clickLog

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ClickLog struct {
	db *sql.DB
}

func NewClickHouseLog(db *sql.DB) IClickLog {
	return &ClickLog{
		db: db,
	}
}

type IClickLog interface {
	LogNewUser(user *domain.UserLog) error
	ShowAllLogs() ([]*domain.UserLog, error)
	Migrate() error
}

func (c *ClickLog) Migrate() error {
	path := filepath.Join("schema", "db", "000001_init_schema.up.sql")
	b, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		return ioErr
	}
	sql := string(b)
	sqlx.MustExec(c.db, sql)
	return nil
}

func (c *ClickLog) LogNewUser(user *domain.UserLog) error {
	query := fmt.Sprintf("INSERT INTO user_logs (id, name, ctreated_at) values ($1, $2, $3)")
	_, err := c.db.Exec(query, user.Id, user.Name, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClickLog) ShowAllLogs() ([]*domain.UserLog, error) {
	var usersLog []*domain.UserLog
	query := fmt.Sprintf("SELECT id, name, created_at FROM user_logs")
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
		//return []domain.UserLog{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var userLog domain.UserLog
		err := rows.Scan(userLog)
		if err != nil {
			return nil, err
			//return []domain.UserLog{}, err
		}
		usersLog = append(usersLog, &userLog)
	}
	if err := rows.Err(); err != nil {
		return nil, err
		//return []domain.UserLog{}, err
	}
	return usersLog, nil
}
