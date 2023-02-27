package product

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Bio struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome product create!")
}

func list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome products list!")
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome product get!")
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome product update!")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome product delete!")
}

func MakeHTTPHandler() http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/product", create).Methods("POST")
	r.HandleFunc("/product", list).Methods("GET")
	r.HandleFunc("/product/:id", get).Methods("GET")
	r.HandleFunc("/product/:id", update).Methods("PUT")
	r.HandleFunc("/product/:id", delete).Methods("DELETE")
	return r
}
