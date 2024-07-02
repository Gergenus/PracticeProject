package storage

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBsource struct {
	db *sql.DB
}

func InitDB() (DBsource, error) {
	db, err := sql.Open("pgx", "postgres://postgres:12345@telegram_db:5432/postgres?sslmode=disable")
	if err != nil {
		return DBsource{}, err
	}
	err = db.Ping()
	if err != nil {
		return DBsource{}, err
	}
	return DBsource{db: db}, nil
}
