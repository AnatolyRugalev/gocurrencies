package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Rates map[string]float64

type Message struct {
	Base  string `json:"base"`
	Rates Rates  `json:"rates"`
}

var apiUrl string

func init() {
	apiUrl = os.Getenv("REMOTE_SVC_URL")
	if apiUrl == "" {
		apiUrl = "https://api.ratesapi.io/api/latest?base=%s"
	}
}

func fetchRates(base string) (Rates, error) {
	resp, err := http.Get(fmt.Sprintf(apiUrl, base))
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("error closing response body: %s", err.Error())
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching rates: status code: %d, response: %s", resp.StatusCode, string(body))
	}
	var msg Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return nil, err
	}
	if msg.Rates == nil {
		return nil, fmt.Errorf("error fetching rates: unsupported response: %s", string(body))
	}
	msg.Rates[msg.Base] = 1.0
	return msg.Rates, nil
}

type Params map[string]string

func (p Params) Str(name string) (string, error) {
	str, ok := p[name]
	if !ok {
		return "", fmt.Errorf("parameter %s is required", name)
	}
	return str, nil
}

func (p Params) Float(name string) (float64, error) {
	str, err := p.Str(name)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(str, 64)
}
