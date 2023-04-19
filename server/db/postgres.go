package db

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	Db *sql.DB
}

func NewPostgres(host, user, password string) (*Postgres, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s sslmode=disable", host, user, password))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Postgres{Db: db}, nil
}