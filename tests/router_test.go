package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/xnpltn/instagram-backend/internal/routes"
)

func SetupAPI(t *testing.T)(string, func()){
	t.Helper()
	server := httptest.NewServer(routes.NewRouter())
	return server.URL, func() {
		server.Close()
	}
}

func TestAPIReadiness(t *testing.T){
	url, close := SetupAPI(t)
	defer close()
	expectedOutput := "API ready\n"

	resp, err := http.Get(url +"/v1/")
	if err!= nil{
		t.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err!= nil{
		t.Fatal(err)
	}
	if !strings.Contains(string(body), expectedOutput){
		t.Fatalf("expected: %q got %q instead", expectedOutput, string(body))
	}
}