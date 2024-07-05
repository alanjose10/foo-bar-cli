package joke

import (
	"encoding/json"
	"io"
	"net/http"
)

const API_URL = "https://icanhazdadjoke.com/"

type JokeResponse struct {
	Id   string `json:"id"`
	Joke string `json:"joke"`
}

func GetRandomJoke() (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var jokeJson JokeResponse
	json.Unmarshal(responseBytes, &jokeJson)
	return jokeJson.Joke, nil
}
