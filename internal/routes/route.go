package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/xnpltn/instagram-backend/internal/handlers"
	"github.com/xnpltn/instagram-backend/internal/applicaction"
)


func NewRouter() *mux.Router{
	mux := mux.NewRouter()
	
	// images
	mux.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads", http.FileServer(http.Dir("uploads/"))))

	// v1
	v1Router := mux.PathPrefix("/v1").Subrouter()
	v1Router.HandleFunc("/", handlers.Readiness)
	app := applicaction.NewApp(v1Router)
	app.LoadRoutes()
	return mux
}


