package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// https://qvault.io/golang/anonymous-structs-golang/
type requestBody struct {
	Status     string `json:"status"`
	Monitoring string `json:"monitoring"`
}

func getApiRequest(url string) requestBody {

	// Method without function

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(bodyBytes))

	params := requestBody{}

	err = json.Unmarshal(bodyBytes, &params)
	if err != nil {
		panic(err)
	}

	// fmt.Println(params.Status)
	// fmt.Println(params.Monitoring)

	return params
	// health := Health{}
	// getAPIContext(url, health.Status)
	// fmt.Println(health.Status)

}
