package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "./var/www/" + pageID + ".html"
	http.ServeFile(w, r, fileName)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler) // ok
	rtr.HandleFunc("/homepage", pageHandler)          // fails
	rtr.HandleFunc("/contact", pageHandler)           // fails
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}
