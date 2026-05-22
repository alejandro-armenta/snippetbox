package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

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

	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet"))
}
