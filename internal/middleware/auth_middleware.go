package middleware

import(
	"net/http"
)

/*
type DBUser struct{
	Name string
	Username string
	Password string
}
type authHandler func(http.ResponseWriter, *http.Request, DBUser)
*/

func AuthMiddleware(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// Auth middleware logic
		handler(w, r)
	}
}