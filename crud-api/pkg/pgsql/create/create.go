package create

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateStructure(db *sql.DB) error {

	// Create the structured table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS articles (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        text TEXT NOT NULL,
        create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`)

	return err
}
