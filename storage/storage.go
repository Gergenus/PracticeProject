package storage

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBsource struct {
	db *sql.DB
}

func InitDB() (DBsource, error) {
	db, err := sql.Open("pgx", "postgres://postgres:1234@localhost:5432/URLsucker")
	if err != nil {
		return DBsource{}, err
	}
	err = db.Ping()
	if err != nil {
		return DBsource{}, err
	}
	return DBsource{db: db}, nil
}
