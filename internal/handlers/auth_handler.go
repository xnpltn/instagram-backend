package handlers

import (
	"net/http"

)

func AuthHandler(w http.ResponseWriter, _ *http.Request){
	w.Write([]byte("Authentication"))
}