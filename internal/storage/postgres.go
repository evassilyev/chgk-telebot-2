package storage

import (
	"context"
	"fmt"
	"github.com/evassilyev/chgk-telebot-2/internal/core/gen/models"
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

// TODO debug and fix it
func (p PostgresStorage) AddGroup(groupId int64) error {
	g := models.Group{
		ID: groupId,
	}
	return g.Insert(context.Background(), p.db)
}

func (p PostgresStorage) GetGroup(groupId int64) (*models.Group, error) {
	return models.GroupByID(context.Background(), p.db, groupId)
}
