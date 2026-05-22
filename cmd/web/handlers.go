package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

	files := []string{
		"../../ui/html/base.tmpl",

		"../../ui/html/partials/nav.tmpl",

		"../../ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//esto escribe eso aqui y lo manda al cliente para que lo vea
	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		//este es para mi
		log.Print(err.Error())
		//este se manda al cliente
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
