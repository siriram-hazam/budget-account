package database

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(cfg *Config) (*Postgres, error) {
	db, err := sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) Close() error {
	return p.db.Close()
}

func (p *Postgres) GetDB() *sql.DB {
	return p.db
}
