package applicaction

import (
	"github.com/xnpltn/instagram-backend/internal/handlers"
	"github.com/xnpltn/instagram-backend/internal/middleware"
)




func(a *App) LoadAuthRoutes() {
	user := handlers.NewUser()
	a.Router.HandleFunc("/auth", handlers.AuthHandler)
	a.Router.HandleFunc("/auth/signup", user.HandlerCreateUser).Methods("POST")
	a.Router.HandleFunc("/auth/login", user.HandlerLoginUser).Methods("POST")
	a.Router.HandleFunc("/auth/{userID}", user.GetUserByID)

}


func (a *App) LoadFollowRoutes(){
	follow := handlers.NewFollow()
	a.Router.HandleFunc("/follow/{userID}", middleware.AuthMiddleware(follow.FollowUser)).Methods("POST")
	a.Router.HandleFunc("/unfollow/{userID}", middleware.AuthMiddleware(follow.UnfollowUser)).Methods("POST")
	a.Router.HandleFunc("/followers", middleware.AuthMiddleware(follow.GetAllFollowing)).Methods("GET")
}


func (a *App)LoadPostsRoutes(){
	post := handlers.NewPost()
	a.Router.HandleFunc("/posts", middleware.AuthMiddleware(post.CreatePosts)).Methods("POST")
	a.Router.HandleFunc("/posts", post.GetPosts).Methods("GET")
	a.Router.HandleFunc("/posts/{postID}", post.GetPostByID).Methods("GET")
	a.Router.HandleFunc("/posts/{postID}", middleware.AuthMiddleware(post.DeletePostByID)).Methods("DELETE")
	a.Router.HandleFunc("/posts/{postID}", middleware.AuthMiddleware(post.EditPostByID)).Methods("PUT")
}


func (a *App) LoadLikeRoutes(){
	like := handlers.NewLike()
	a.Router.HandleFunc("/like/{postID}", middleware.AuthMiddleware(like.LikePostByID)).Methods("POST")
	a.Router.HandleFunc("/unlike/{postID}", middleware.AuthMiddleware(like.UnikePostByID)).Methods("POST")
	a.Router.HandleFunc("/likes/{postID}", like.GetLikesCountByID).Methods("GET")
}