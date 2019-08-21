// gocurrencies is REST api service for currency pairs of USD/EUR and other
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Env() string {
	return os.Getenv("REMOTE_SVC_URL")
}

func main() {
	remoteSvc := Env()
	router := mux.NewRouter()
	router.HandleFunc("/currencies/{"+baseParameter+":[A-Z]+}", func(w http.ResponseWriter, r *http.Request) {
		AttitudeBase(remoteSvc, w, r)
	}).Methods("GET")
	router.HandleFunc("/currencies/{"+baseParameter+":[A-Z]+}/{"+targetParameter+":[A-Z]+}", func(w http.ResponseWriter, r *http.Request) {
		AttitudePair(remoteSvc, w, r)
	}).Methods("GET")
	router.HandleFunc("/currencies/{"+baseParameter+":[A-Z]+}/{"+targetParameter+":[A-Z]+}/{"+sumParameter+":[0-9\\.]+}", func(w http.ResponseWriter, r *http.Request) {
		AttitudeSum(remoteSvc, w, r)
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
