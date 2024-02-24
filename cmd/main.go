package main

import (
	"context"
	"log"
	"os"

	"os/signal"

	_ "github.com/lib/pq"
	"github.com/xnpltn/instagram-backend/internal/applicaction"
	"github.com/xnpltn/instagram-backend/internal/routes"
)

func main(){
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	app := applicaction.NewApp(routes.NewRouter())
	log.Fatal(app.Start(ctx))
}