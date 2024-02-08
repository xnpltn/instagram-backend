package utils


import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string){
	if code > 499 {
		log.Println("Responding with 500 err: ", message)
	}
	type errResponce struct{
		Error string `json:"error"`
	}
	RespondWithJson(w, code, errResponce{
		Error: message,
	})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err != nil{
		w.WriteHeader(500)
		log.Printf("Failed to marshal JSON responce %v", payload)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}