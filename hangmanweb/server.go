package main

import (
	"fmt"
	"hangmanweb/utils"
	"net/http"
	"os"
	"text/template"
	"golang.org/x/tools/godoc/util"
)

const port = ":8080"

func StartGame() Utils.HangManData {
	data := utils.HangManData{Attempts: 10, Error: ""}
	data.Word = util.ReadFileName(os.Args[1])
	data.ToFind =
	data.HangmanPositions = utils.Drawhangman(./hangman.txt)
}

/// modif endessous ///
func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "")
	l := r.FormValue("name")
	if l != "" {
		hangmanweb.HangMan(data, l)
	}
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
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("[ERROR] - Server could not start properly.\n ", err)
	}
}

