package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const port = ":8080"

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Index)
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))
	fmt.Println("(http://localhost:8080/) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
