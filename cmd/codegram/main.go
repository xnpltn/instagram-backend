package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/xnpltn/codegram/internal/routes"
)


func main(){
	server := &http.Server{
		Handler: routes.NewRouter(),
		Addr: "localhost:9090",
	}

	 

	fmt.Println("server starting http://localhost:9090")

	err :=server.ListenAndServe()
	if err != nil {
		log.Fatal("error starting the server")
	}
}