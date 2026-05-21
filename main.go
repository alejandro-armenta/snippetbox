package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	a := r.PathValue("id")

	id, error := strconv.Atoi(a)

	if error != nil || id < 1 {

		http.NotFound(
			w,
			r)

		return
	}

	msg := fmt.Sprintf(
		"Display a specific snippet with ID %d",
		id)

	w.Write(
		[]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func main() {
	//router
	/*
		eachtime a Request is made the server gets it and it passes it onto the mux
		the mux controls to which handler it gives the request based on url
		and dispatch the request to the matching handler.
	*/

	mux := http.NewServeMux()

	//catch all
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}
