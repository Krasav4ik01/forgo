package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	Avg_grades, Happines float64
}

func (u *User) setName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	object2 := User{"Bob", 28, 50, 4.2, 0.8}
	// object2.setName("Alan")
	// fmt.Fprintf(w, "User name is: "+object2.name)
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, object2)
}
func second_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the second page")
}

func fetchRequests() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/second/", second_page)
	http.ListenAndServe(":8080", nil)
}

func main() {

	fetchRequests()
}
