package applicaction

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/xnpltn/instagram-backend/internal/database"
)


type App struct{
	DB *sql.DB
	Router *mux.Router
}

  

func NewApp( router *mux.Router) *App{
	return &App{
		Router: router,
	}
}

func(a *App)Start(ctx context.Context) error {
	db, err := database.Connect()
	if err!= nil{
		return err
	}
	a.DB = db
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	s := &http.Server{
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(a.Router),
		Addr: ":8080",
	}
	errorChan := make(chan error, 1)
	go func() {
		err := s.ListenAndServe()
		if err!= nil{
			errorChan <- err
		}
		close(errorChan)
	}()
	err = <- errorChan

	select {
	case <-errorChan:
		return err
	case <- ctx.Done():
		fmt.Println("shouting down....")
		timeOutContext, cancel := context.WithTimeout(context.Background(), time.Second *10)
		defer cancel()
		err := s.Shutdown(timeOutContext)
		if err!= nil{
			log.Printf("error shuting down: %s", err.Error())
			os.Exit(1)
		}
		os.Exit(1)
		return nil
	}
	
}


func (a *App)LoadRoutes(){
	a.LoadAuthRoutes()
	a.LoadFollowRoutes()
	a.LoadLikeRoutes()
	a.LoadPostsRoutes()
}