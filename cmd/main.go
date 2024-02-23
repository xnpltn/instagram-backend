package main

import (
	"log"
	_ "github.com/lib/pq"
	"github.com/xnpltn/instagram-backend/internal/applicaction"
	"github.com/xnpltn/instagram-backend/internal/routes"
)

func main(){
	app := applicaction.NewApp(routes.NewRouter())
	log.Fatal(app.Start())
}