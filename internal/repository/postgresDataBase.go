package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host    string
	Port    string
	Dbname  string
	Sslmode string
}

func NewConnectPostgres(cfg Config) (*sqlx.DB, error) {
	connect := fmt.Sprintf("host=%v port=%v dbname=%v sslmode=%v", cfg.Host, cfg.Port, cfg.Dbname, cfg.Sslmode)
	db, err := sqlx.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
