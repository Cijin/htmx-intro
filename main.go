package main

import (
	"fmt"
	"htmx-intro/templates"
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	t := templates.New("public/views/*.html")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Render(w, "home", nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.HandleFunc("/mouseEnter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Recieved")
		return
	})

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		someItems := []string{"yes", "no", "maybe"}
		res := strings.Join(someItems, "\n")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
		return
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
