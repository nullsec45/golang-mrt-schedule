package client

import (
	"net/http"
	"errors"
	"io/ioutil"
)

func DoRequest(client *http.Client, url string)([]byte, error){
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Unexpected status code: " + response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}