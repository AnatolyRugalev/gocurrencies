package main

import (
	"log"
	"strconv"
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
	currencies, err := currencyValues(remoteSvc)
	if err != nil {
		log.Printf("%s [AttitudePair] getting currency values failed. Debug: %v", r.RemoteAddr, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	baseKey, err := parameter(baseParameter, r)
	if err != nil {
		log.Printf("%s [AttitudePair] missing parameter `%s` in http.Request. Debug: %v", r.RemoteAddr, baseParameter, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	baseValue, ok := currencies[baseKey]
	if !ok {
		log.Printf("%s [AttitudePair] unsupported base currency: %s", r.RemoteAddr, baseKey)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	targetKey, err := parameter(targetParameter, r)
	if err != nil {
		log.Printf("%s [AttitudePair] missing parameter `%s` in http.Request. Debug: %v", r.RemoteAddr, targetParameter, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	targetValue, ok := currencies[targetKey]
	if !ok {
		log.Printf("%s [AttitudePair] unsupported target currency: %s", r.RemoteAddr, targetKey)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	output := targetValue / baseValue
	outputJson, _ := json.Marshal(output)
	w.Write([]byte(string(outputJson)))
}

func AttitudeSum(remoteSvc string, w http.ResponseWriter, r *http.Request) {
	currencies, err := currencyValues(remoteSvc)
	if err != nil {
		log.Printf("%s [AttitudeSum] getting currency values failed. Debug: %v", r.RemoteAddr, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	baseKey, err := parameter(baseParameter, r)
	if err != nil {
		log.Printf("%s [AttitudeSum] missing parameter `%s` in http.Request. Debug: %v", r.RemoteAddr, baseParameter, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	baseValue, ok := currencies[baseKey]
	if !ok {
		log.Printf("%s [AttitudeSum] unsupported base currency: %s", r.RemoteAddr, baseKey)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	targetKey, err := parameter(targetParameter, r)
	if err != nil {
		log.Printf("%s [AttitudeSum] missing parameter `%s` in http.Request. Debug: %v", r.RemoteAddr, targetParameter, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	targetValue, ok := currencies[targetKey]
	if !ok {
		log.Printf("%s [AttitudeSum] unsupported target currency: %s", r.RemoteAddr, targetKey)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sumStr, err := parameter(sumParameter, r)
	if err != nil {
		log.Printf("%s [AttitudeSum] missing parameter `%s` in http.Request. Debug: %v", r.RemoteAddr, sumParameter, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sumValue, err := strconv.ParseFloat(sumStr, 32)
	if err != nil {
		log.Printf("%s [AttitudeSum] invalid sum parameter: %s", r.RemoteAddr, sumStr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	output := (targetValue / baseValue) * float32(sumValue)
	outputJson, _ := json.Marshal(output)
	w.Write([]byte(string(outputJson)))
}
