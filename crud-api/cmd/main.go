package main

import (
	// castome class
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	// my imports
	"crud-api.com/pkg/pgsql/connect"
	"crud-api.com/pkg/pgsql/create"
	"crud-api.com/pkg/service/articles/createArticle"
	"crud-api.com/pkg/service/articles/getArticle"
	"crud-api.com/pkg/service/articles/updateArticle"
	"crud-api.com/pkg/service/articles/deleteArticle"
)

func main() {
	HandleFunc()
}

func HandleFunc() {
	// ConnectAndCreateStructureDb() // use one time when comment
	route := mux.NewRouter()
	route.HandleFunc("/api/v1/create", createArticle.CreateArticle).Methods("POST")
	route.HandleFunc("/api/v1/article/{id}", getArticle.GetDataByID).Methods("GET")
	route.HandleFunc("/api/v1/article/{id}", updateArticla.UpdateArticle).Methods("PATCH")
	route.HandleFunc("/api/v1/article/delete/{id}", deleteArticle.DeleteArticle).Methods("DELETE")

	http.Handle("/", route)
	http.ListenAndServe(":8080", nil)
}

func ConnectAndCreateStructureDb() {
	db, err := connect.ConnectToDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = create.CreateStructure(db)
	if err != nil {
		panic(err)
	}
}
