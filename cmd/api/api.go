package main

import (
	c "github.com/Gusarov2k/library_io/internal/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test", c.GetJSON).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
