package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from alex"))
}

func main() {
	//router
	/*
		eachtime a Request is made the server gets it and it passes it onto the mux
		the mux controls to which handler it gives the request based on url
		and dispatch the request to the matching handler.
	*/

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}
