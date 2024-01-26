package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/xnpltn/codegram/internal/routes"
	_ "github.com/lib/pq"
	"github.com/xnpltn/codegram/internal/models"
)




func main(){

	db, err := models.InitDB()
	if err != nil{
		log.Fatal("failed to initialize db lol")
	}

	fmt.Println(db)
	server := &http.Server{
		Handler: routes.NewRouter(),
		Addr: "localhost:9090",
	}
	
	fmt.Println("server starting http://localhost:9090")

	err =server.ListenAndServe()
	if err != nil {
		log.Fatal("error starting the server")
	}
}