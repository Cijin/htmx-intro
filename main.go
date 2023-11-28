package main

import (
	"htmx-intro/templates"
	"log"
	"net/http"
)

type BarData struct {
	Trigger string
	Width   int
}

type Data struct {
	Status string
	Bar    BarData
}

var (
	data     Data
	progress int
)

func main() {
	mux := http.NewServeMux()
	t := templates.New("public/views/*.html")
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/public/", http.StripPrefix("/public/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Render(w, "home", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		progress = 0
		data = Data{
			"Running",
			BarData{"every 600ms", progress},
		}

		err := t.Render(w, "progress", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/job", func(w http.ResponseWriter, r *http.Request) {
		data = Data{
			"Complete",
			BarData{"none", progress},
		}
		err := t.Render(w, "progress", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/job/progress", func(w http.ResponseWriter, r *http.Request) {
		progress += 20
		barData := BarData{data.Bar.Trigger, progress}

		if progress > 100 {
			w.Header().Add("HX-Trigger", "done")
		}

		err := t.Render(w, "bar", barData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
