package joke

import (
	"io"
	"net/http"
)

const API_URL = "https://icanhazdadjoke.com/"

func GetRandomJoke() (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "text/plain")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	responseString := string(responseBytes)
	return responseString, nil

}
