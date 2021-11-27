package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ipReponse struct {
	IP string `json:"ip"`
}

func GetMyIP() (string, error) {
	req, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	fmt.Print(string(body))

	var ipResponse ipReponse
	err = json.Unmarshal(body, &ipResponse)
	if err != nil {
		return "", err
	}
	return ipResponse.IP, nil
}
