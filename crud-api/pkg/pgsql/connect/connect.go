package connect

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5431
	user     = "golang"
	password = "password"
	dbname   = "golang_db"
)

func ConnectToDb() (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		return nil, err
	}
	return db, nil
}
