package updateArticla

import (
	// "encoding/json"
	// "fmt"
	"net/http"

	_ "github.com/lib/pq"

	// my imports
	"crud-api.com/pkg/pgsql/connect"
)

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	db, err := connect.ConnectToDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверка метода запроса
	if r.Method != http.MethodPatch {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Извлечение параметров title и text из тела запроса
	title := r.FormValue("title")
	text := r.FormValue("text")


	// update
	sqlStatement := `
	UPDATE articles
	SET title = $2, text = $3
	WHERE id = $1;`
	_, err = db.Exec(sqlStatement, 2, title, text)

	if err != nil {
		panic(err)
	}

}
