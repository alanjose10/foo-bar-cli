package ip

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMyIp(t *testing.T) {
	mockResponse := `{"ip": "12.34.56.78"}`
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	ip, err := GetMyIp(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedIp := "12.34.56.78"

	if ip != expectedIp {
		t.Errorf("expected %v, got %v", expectedIp, ip)
	}

}
