package domain

import (
	"bytes"
	"encoding/gob"
	"time"

	lg "github.com/cookienyancloud/testrpckafkapsqlclick/pkg/logger/logger"
)

type User struct {
	Id           string `db:"id" redis:"id"`
	Name         string `db:"name" redis:"name"`
	PasswordHash string `db:"password_hash" redis:"password_hash"`
}

type UserLog struct {
	Id        string    `db:"id" redis:"id"`
	Name      string    `db:"name" redis:"name"`
	Key       string    `db:"key" redis:"key"`
	CreatedAt time.Time `db:"created_at" redis:"created_at"`
}

type Encoder interface {
	Encode() ([]byte, error)
	Length() int
}

func (u UserLog) Encode() ([]byte, error) {
	var msgBytes bytes.Buffer
	enc := gob.NewEncoder(&msgBytes)
	err := enc.Encode(UserLog{
		Id:        u.Id,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
	})
	return msgBytes.Bytes(), err
}

func (u UserLog) Length() int {
	var msgBytes bytes.Buffer
	enc := gob.NewEncoder(&msgBytes)
	err := enc.Encode(UserLog{
		Id:        u.Id,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
	})
	if err != nil {
		lg.Errorf("error getting length:%v", err)
		return 0
	}
	return len(msgBytes.Bytes())
}

func Decode(msgByte *bytes.Buffer, msg *UserLog) error {
	dec := gob.NewDecoder(msgByte)
	err := dec.Decode(&msg)
	return err
}
