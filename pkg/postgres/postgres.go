package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vgekko/go-tasks-graphql/config"
)

func NewPostgres(cfg config.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
