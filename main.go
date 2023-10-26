package main

import (
	"html/template"
	"log"
	"net/http"
)

type Book struct {
	Title  string
	Author string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		books := map[string][]Book{
			"Books": {
				{Title: "Fun Times", Author: "Michael"},
				{Title: "Deep Work", Author: "Cal"},
				{Title: "ABC", Author: "Ken"},
			},
		}
		tmpl.Execute(w, books)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.ExecuteTemplate(w, "book-list-element", Book{Title: title, Author: author})
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-book/", h2)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
