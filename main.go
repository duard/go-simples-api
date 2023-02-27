package main

import (
	"fmt"

	"github.com/charmbracelet/log"

	brand "github.com/duard/go-simples-api/domains/brands"

	"net/http"

	product "github.com/duard/go-simples-api/domains/products"

	"github.com/gorilla/mux"
)

func init() {
	fmt.Print(`initializating...`)
}
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome home!")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	brandHandler := brand.MakeHTTPHandler()
	router.Handle("/brand", brandHandler)
	router.Handle("/brand/", brandHandler)

	productHandler := product.MakeHTTPHandler()
	router.Handle("/product", productHandler)
	router.Handle("/product/", productHandler)

	log.Info("Server is running at http://localhost:8080 🍪")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("There's an error with the server", err)
	}
}
