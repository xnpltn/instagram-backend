package middleware

import (
	"fmt"
	"net/http"

	"github.com/xnpltn/instagram-backend/internal/models"
	"github.com/xnpltn/instagram-backend/internal/utils"
)

type authHandler func(http.ResponseWriter, *http.Request, models.DBUser)

func AuthMiddleware(handler authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()
		tokenKey := ""
		for _, cookie := range cookies{
			if cookie.Name == "Token"{
				tokenKey += cookie.Value
			}
		}


		if tokenKey == ""{
			utils.RespondWithError(w, http.StatusUnauthorized, "You need to login")
			return
		}		
		user, err := utils.VerifyJWT(tokenKey)
		if err != nil{
			utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Could not get user: %v", err))
			return
		}
		
		handler(w, r, user)
	}
}