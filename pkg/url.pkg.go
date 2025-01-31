package pkg

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func CallURLGet(url string) (result string, err error) {
	log.Println(url)

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close() // Ensure the body is closed after reading

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Convert the body to a string and return it
	result = string(body)
	log.Println(result)

	// Check if the status code is not 2xx (success)
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", errors.New(result)
	}
	return result, nil
}
