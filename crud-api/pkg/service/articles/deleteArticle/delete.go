package deleteArticle

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	// my imports
	"crud-api.com/pkg/pgsql/connect"
)

func DeleteArticle(w http.ResponseWriter, r *http.Request) {

	// get parametr id
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "User ID: %s\n", id)

	db, err := connect.ConnectToDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверка метода запроса
	if r.Method != http.MethodDelete {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// delete
	deleteStmt := `DELETE FROM "articles" where id=$1`
	_, e := db.Exec(deleteStmt, id)

	if e != nil {
		panic(err)
	}

}
