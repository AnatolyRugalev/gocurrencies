package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Currencyes map[string]float32

type Message struct {
	Base  string     `json:"base"`
	Rates Currencyes `json:"rates"`
}

func currencyValues(remoteSvc string) (result Currencyes, err error) {
	resp, err := http.Get(remoteSvc)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var msg Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return
	}

	result = Currencyes{msg.Base: 1.0}
	for currency, attitude := range msg.Rates {
		result[currency] = attitude
	}

	return
}

func attitudeByBase(baseKey string, currencies Currencyes) (result Currencyes, err error) {
	baseValue, ok := currencies[baseKey]
	if !ok {
		return result, errors.New("missing base")
	}
	result = Currencyes{}
	for currency, attitude := range currencies {
		if currency != baseKey {
			result[currency] = attitude/baseValue
		}
	}
	return
}

func parameter(parameterKey string, r *http.Request) (string, error) {
	params := mux.Vars(r)
	result, ok := params[parameterKey]
	if !ok {
		return "", errors.New("missing parameter")
	}
	return result, nil
}
