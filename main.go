// gocurrencies is REST api service for currency pairs of USD/EUR and other
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	route := fmt.Sprintf("/currencies/{%s:[A-Z]+}", baseParameter)
	router.HandleFunc(route, Handler(GetRates)).Methods("GET")

	route = fmt.Sprintf("%s/{%s:[A-Z]+}", route, targetParameter)
	router.HandleFunc(route, Handler(GetRate)).Methods("GET")

	route = fmt.Sprintf("%s/{%s:[0-9\\.]+}", route, sumParameter)
	router.HandleFunc(route, Handler(CalculateSum)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
