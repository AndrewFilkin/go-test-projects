package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5431
	user     = "golang"
	password = "password"
	dbname   = "golang_db"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func SaveArticle(w http.ResponseWriter, r *http.Request) {

	//get data with form
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusInternalServerError)
        return
    }

    title := r.Form.Get("title")
    text := r.Form.Get("text")

	pgsqlConnect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	// open database
	db, err := sql.Open("postgres", pgsqlConnect)
	CheckError(err)

	defer db.Close()

	// add data
	insert, err := db.Query(fmt.Sprintf("INSERT INTO articles (title, text) VALUES('%s', '%s')", title, text))
	CheckError(err)

	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func HandleFunc() {
	route := mux.NewRouter()
	route.HandleFunc("/", Index).Methods("GET")
	route.HandleFunc("/create/", Create).Methods("GET")
	route.HandleFunc("/save_article/", SaveArticle).Methods("POST")
	http.Handle("/", route)
	
	//for add css file
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	HandleFunc()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
