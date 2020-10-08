package storage

import (
	"context"
	"time"

	"github.com/BarTar213/notificator/config"
	"github.com/BarTar213/notificator/models"
	"github.com/go-pg/pg/v10"
)

const (
	all    = "*"
)

type Postgres struct {
	db *pg.DB
}

type Storage interface {
	AddTemplate(template *models.Template) error
}

func NewPostgres(config *config.Postgres) (Storage, error) {
	db := pg.Connect(&pg.Options{
		Addr:        config.Address,
		User:        config.User,
		Password:    config.Password,
		Database:    config.Database,
		DialTimeout: 5 * time.Second,
	})

	err := db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}
