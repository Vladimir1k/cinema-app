package repository

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func OpenDB() (*sql.DB, error) {
	var cfg = Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "admin",
		Password: "root",
		DBName:   "db_cinema",
		SSLMode:  "disable",
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
