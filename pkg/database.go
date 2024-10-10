package pkg

import (
	"database/sql"

	_ "github.com/lib/pq" // driver do PostgreSQL
)

// Connect conecta ao banco de dados PostgreSQL
func Connect(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
