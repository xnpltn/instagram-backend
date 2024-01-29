package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xnpltn/codegram/internal/handlers"
	"github.com/xnpltn/codegram/internal/middleware"
)


func NewRouter() *mux.Router{
	mux := mux.NewRouter()
	mux.HandleFunc("/", testHomeHandler)
	v1Router := mux.PathPrefix("/v1").Subrouter()
	v1Router.HandleFunc("", v1Handler)

	user := handlers.NewUser()
	v1Router.HandleFunc("/auth", handlers.AuthHandler)
	v1Router.HandleFunc("/auth/signup", user.HandlerCreateUser).Methods("POST")
	v1Router.HandleFunc("/auth/login", user.HandlerLoginUser).Methods("POST")

	post := handlers.NewPost()
	v1Router.HandleFunc("/posts", middleware.AuthMiddleware(post.CreatePosts)).Methods("POST")
	v1Router.HandleFunc("/posts", post.GetPosts).Methods("GET")
	v1Router.HandleFunc("/posts/{postID}", post.GetPostByID).Methods("GET")
	v1Router.HandleFunc("/posts/{postID}", middleware.AuthMiddleware(post.DeletePostByID)).Methods("DELETE")

	return mux
}


func testHomeHandler(w http.ResponseWriter, _ *http.Request){
	w.Write([]byte("It works"))
}

func v1Handler(w http.ResponseWriter, _ *http.Request){
	w.Write([]byte("It works on v1"))
}