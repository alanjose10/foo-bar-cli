package ip

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IpApiResponse struct {
	Ip string `json:"ip"`
}

const API_URL = "https://api.ipify.org"

func GetMyIp() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s?format=json", API_URL))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to invoke API: %s", resp.Status)
	}

	var ipApiResponse IpApiResponse

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&ipApiResponse)
	if err != nil {
		return "", err
	}
	return ipApiResponse.Ip, nil
}
