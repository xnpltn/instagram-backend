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
	expectedOutput :=struct{
		Body string
		Code int
	}{
		Body: "API ready\n",
		Code: 200,
	}

	resp, err := http.Get(url +"/v1/")
	if err!= nil{
		t.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err!= nil{
		t.Fatal(err)
	}
	if resp.StatusCode != expectedOutput.Code{
		t.Errorf("expected: %q got %q instead", expectedOutput.Code, resp.StatusCode)
	}
	if !strings.Contains(string(body), expectedOutput.Body){
		t.Fatalf("expected: %q got %q instead", expectedOutput.Body, string(body))
	}
}