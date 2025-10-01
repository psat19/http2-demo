package main

import (
	"bytes"
	"log"
	"net/http"
	"text/template"
)

func serveHTML(w http.ResponseWriter, r *http.Request) {
	render(w)
}

func render(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("./ui/html/index.tmpl"))

	buf := new(bytes.Buffer)

	err := tmpl.Execute(buf, nil)
	if err != nil {
		log.Printf("error executing the template set")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", serveHTML)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

func main() {

	srv := http.Server{
		Addr:    ":8081",
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
