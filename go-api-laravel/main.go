package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func(u *User) SetNewName(newName string) {
	u.Name = newName
}

func (u *User) GetAllInfo() string {
	return fmt.Sprintf("%s is %d years old and have %d money", u.Name, u.Age, u.Money)
}

func main() {
	hendleRequest()
}

func hendleRequest() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/contacts/", ContactsPage)
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	Bob := User{
		Name:       "Bob",
		Age:        28,
		Money:      4000,
		Avg_grades: 95.3,
		Happiness:  99.9,
		Hobbies: []string {"qwer", "asdf"},
	}

	Bob.SetNewName("Andrew")
	// fmt.Fprintf(w, `<b>Main Text </b>`)
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, Bob)

}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts")
}
