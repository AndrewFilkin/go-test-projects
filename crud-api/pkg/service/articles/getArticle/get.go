package getArticle

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	// my imports
	"crud-api.com/pkg/pgsql/connect"
)

type Response struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func GetDataByID(w http.ResponseWriter, r *http.Request) {
	// get parametr id
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "User ID: %s\n", id)

	// open database
	db, err := connect.ConnectToDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверка метода запроса
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query(fmt.Sprintf(`SELECT "title", "text" FROM "articles" WHERE "id" = %s`, id))
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var title string
	var text string

	for rows.Next() {

		err = rows.Scan(&title, &text)
		if err != nil {
			panic(err)
		}

		if len(title) == 0 {
			fmt.Fprintf(w, "Emptu")
		}

		// fmt.Fprintf(w, "Title: %s Text: %s", title, text)
	}

	// Создание структуры Response
	response := Response{
		Title: title,
		Text:  text,
	}

	// Преобразование в JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Ошибка преобразования в JSON", http.StatusInternalServerError)
		return
	}

	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Возврат JSON-ответа
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
