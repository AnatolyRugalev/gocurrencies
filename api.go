package main

import (
	"log"
	"net/http"
	"encoding/json"
)

const (
	baseParameter   = "base"
	targetParameter = "target"
	sumParameter    = "sum"
)

func AttitudeBase(remoteSvc string, w http.ResponseWriter, r *http.Request) {
	currencies, err := currencyValues(remoteSvc)
	if err != nil {
		log.Printf("%s [AttitudeBase] getting currency values failed. Debug: %v", r.RemoteAddr, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	baseKey, err := parameter(baseParameter, r)
	if err != nil {
		log.Printf("%s [AttitudeBase] missing parameter `%s` in http.Request. Debug: %v", r.RemoteAddr, baseParameter, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, ok := currencies[baseKey]; !ok {
		log.Printf("%s [AttitudeBase] unsupported base currency: %s", r.RemoteAddr, baseKey)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	output, err := attitudeByBase(baseKey, currencies)
	if err != nil {
		log.Printf("%s [AttitudeBase] unexpected behavior. Debug: %v", r.RemoteAddr, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	outputJson, _ := json.Marshal(output)
	w.Write([]byte(string(outputJson)))
}

func AttitudePair(remoteSvc string, w http.ResponseWriter, r *http.Request) {
	AttitudeBase(remoteSvc, w, r)
}

func AttitudeSum(remoteSvc string, w http.ResponseWriter, r *http.Request) {
	AttitudeBase(remoteSvc, w, r)
}
