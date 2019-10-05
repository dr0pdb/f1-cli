package network

import (
	"io/ioutil"
	"log"
	"net/http"
)

// MakeGetRequest makes a HTTP GET request taking in the URL as the parameter.
// The response body is returned as a byte slice.
func MakeGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return []byte{}, err
	}

	return body, nil
}
