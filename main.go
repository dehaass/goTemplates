package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/header.tmpl", "templates/footer.tmpl", "keyword.tmpl"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "keyword", nil)
	})
	http.ListenAndServe(":8080", nil)
}
