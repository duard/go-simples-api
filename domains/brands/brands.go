package brand

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
	fmt.Fprintf(w, "Welcome brand create!")
}

func list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome brands list!")
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome brand get!")
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome brand update!")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome brand delete!")
}

func MakeHTTPHandler() http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/brand", create).Methods("POST")
	r.HandleFunc("/brand", list).Methods("GET")
	r.HandleFunc("/brand/:id", get).Methods("GET")
	r.HandleFunc("/brand/:id", update).Methods("PUT")
	r.HandleFunc("/brand/:id", delete).Methods("DELETE")
	return r
}
