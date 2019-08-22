package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HandlerFunc func(params Params) (interface{}, error)

const (
	baseParameter   = "base"
	targetParameter = "target"
	sumParameter    = "sum"
)

func Handler(handler HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		out, err := handler(mux.Vars(r))
		if err != nil {
			handleError(w, err)
		} else {
			w.WriteHeader(http.StatusOK)
			output(w, out)
		}
	}
}

func output(w http.ResponseWriter, out interface{}) {
	w.Header().Set("content-type", "application/json")
	j, err := json.Marshal(out)
	if err != nil {
		log.Fatalf("error marshalling output")
	}
	_, err = w.Write(j)
	if err != nil {
		log.Fatalf("error writing output")
	}
}

func handleError(w http.ResponseWriter, err error) {
	log.Printf("error handling request: %s", err.Error())
	jsonError := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}
	w.WriteHeader(http.StatusInternalServerError)
	output(w, jsonError)
}

// GetRates GET /currencies/<base> API method handler
func GetRates(params Params) (interface{}, error) {
	return getRates(params)
}

func getRates(params Params) (Rates, error) {
	base, err := params.Str(baseParameter)
	if err != nil {
		return nil, fmt.Errorf("missing base parameter: %s", err.Error())
	}
	return fetchRates(base)
}

// GetRate GET /currencies/<base>/<target> API method handler
func GetRate(params Params) (interface{}, error) {
	return getRate(params)
}

func getRate(params Params) (float64, error) {
	rates, err := getRates(params)
	if err != nil {
		return 0.0, fmt.Errorf("error getting rates: %s", err.Error())
	}
	target, err := params.Str(targetParameter)
	if err != nil {
		return 0.0, fmt.Errorf("missing target parameter: %s", err.Error())
	}
	value, ok := rates[target]
	if !ok {
		return 0.0, fmt.Errorf("unsupported target currency: %s", target)
	}
	return value, nil
}

// CalculateSum GET /currencies/<base>/<target>/<sum> API method handler
func CalculateSum(params Params) (interface{}, error) {
	rate, err := getRate(params)
	if err != nil {
		return nil, fmt.Errorf("error getting rate: %s", err.Error())
	}
	sum, err := params.Float(sumParameter)
	if err != nil {
		return nil, fmt.Errorf("missing sum parameter: %s", err.Error())
	}
	return rate * sum, nil
}
