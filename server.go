package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "xHello word!")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A propos de moi:")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("(http://localhost:8080/) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
