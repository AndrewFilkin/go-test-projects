package createArticle

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
	// my imports
	"crud-api.com/pkg/pgsql/connect"
)

type Response struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	db, err := connect.ConnectToDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверка метода запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Извлечение параметров title и text из тела запроса
	title := r.FormValue("title")
	text := r.FormValue("text")

	// Проверка наличия параметров
	if title == "" || text == "" {
		http.Error(w, "Отсутствуют обязательные параметры", http.StatusBadRequest)
		return
	}

	// write to database
	insert, err := db.Query(fmt.Sprintf("INSERT INTO articles (title, text) VALUES('%s', '%s')", title, text))

	if err != nil {
		panic(err)
	}

	defer insert.Close()

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
