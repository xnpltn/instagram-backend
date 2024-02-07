package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/xnpltn/codegram/internal/routes"
	_ "github.com/lib/pq"
	"github.com/gorilla/handlers"
	"github.com/xnpltn/codegram/internal/database"
)




func main(){

	_, err := database.Connect()
	if err != nil{
		log.Fatal("failed to initialize db lol")
	}

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	server := &http.Server{
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(routes.NewRouter()),
		Addr: "localhost:9090",
	}
	
	fmt.Println("server starting http://localhost:9090")
	
	err =server.ListenAndServe()
	if err != nil {
		log.Fatal("error starting the server", err)
	}
}