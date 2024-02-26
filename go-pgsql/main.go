package main

import (
	"fmt"
	"database/sql"
	"github.com/lib/pq"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

const (
	host     = "localhost"
	port     = 5431
	user     = "golang"
	password = "password"
	dbname   = "golang_db"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// add data
	// insert, err := db.Query("INSERT INTO users (name, age) VALUES('Bob', 35)")
	// CheckError(err)
	// defer insert.Close()

	// Select Data
	res, err := db.Query("SELECT name, age FROM users")
	CheckError(err)

	for res.Next() {
		var user User
		err := res.Scan(&user.Name, &user.Age)
		CheckError(err)
		fmt.Printf("User %s with age %d\n", user.Name, user.Age)
	}

	defer res.Close()



}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
