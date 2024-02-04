package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/xnpltn/codegram/internal/handlers"
	"github.com/xnpltn/codegram/internal/middleware"
)


func NewRouter() *mux.Router{
	mux := mux.NewRouter()
	
	// images
	mux.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads", http.FileServer(http.Dir("uploads/"))))

	// v1
	v1Router := mux.PathPrefix("/v1").Subrouter()

	// users, auth
	user := handlers.NewUser()
	v1Router.HandleFunc("/auth", handlers.AuthHandler)
	v1Router.HandleFunc("/auth/signup", user.HandlerCreateUser).Methods("POST")
	v1Router.HandleFunc("/auth/login", user.HandlerLoginUser).Methods("POST")
	v1Router.HandleFunc("/users/{userID}", user.GetUserByID)

	// follow
	follow := handlers.NewFollow()
	v1Router.HandleFunc("/follow/{userID}", middleware.AuthMiddleware(follow.FollowUser)).Methods("POST")
	v1Router.HandleFunc("/unfollow/{userID}", middleware.AuthMiddleware(follow.UnfollowUser)).Methods("POST")
	v1Router.HandleFunc("/followers", middleware.AuthMiddleware(follow.GetAllFollowing)).Methods("GET")

	// posts
	post := handlers.NewPost()
	v1Router.HandleFunc("/posts", middleware.AuthMiddleware(post.CreatePosts)).Methods("POST")
	v1Router.HandleFunc("/posts", post.GetPosts).Methods("GET")
	v1Router.HandleFunc("/posts/{postID}", post.GetPostByID).Methods("GET")
	v1Router.HandleFunc("/posts/{postID}", middleware.AuthMiddleware(post.DeletePostByID)).Methods("DELETE")
	v1Router.HandleFunc("/posts/{postID}", middleware.AuthMiddleware(post.EditPostByID)).Methods("PUT")

	return mux
}



