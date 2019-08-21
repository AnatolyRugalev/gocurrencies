package main

import (
	"io/ioutil"
	"net/http"
)

const (
	baseParameter   = "base"
	targetParameter = "target"
	sumParameter    = "sum"
)

func AttitudeBase(remoteSvc string, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	w.Write(body)
}

func AttitudePair(remoteSvc string, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	w.Write(body)
}

func AttitudeSum(remoteSvc string, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	w.Write(body)
}
