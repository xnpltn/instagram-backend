package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/xnpltn/codegram/internal/handlers"
)


func NewRouter() *mux.Router{
	mux := mux.NewRouter()
	mux.HandleFunc("/", testHomeHandler)
	v1Router := mux.PathPrefix("/v1").Subrouter()
	v1Router.HandleFunc("/", v1Handler)
	v1Router.HandleFunc("/auth", handlers.AuthHandler)

	return mux
}


func testHomeHandler(w http.ResponseWriter, _ *http.Request){
	w.Write([]byte("It works"))
}

func v1Handler(w http.ResponseWriter, _ *http.Request){
	w.Write([]byte("It works on v1"))
}