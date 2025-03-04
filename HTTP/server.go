package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	const port string = ":8080"
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running!")
	})
	log.Println("Listening on " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
