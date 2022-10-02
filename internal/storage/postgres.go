package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage(user, pass, host, db string, port uint16) (*PostgresStorage, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, db)
	dbd, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		db: dbd,
	}, nil
}

func (p PostgresStorage) AddGroup(chatId int64) error {
	return nil
}
