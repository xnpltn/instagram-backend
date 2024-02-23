package middleware

import (
	"fmt"
	"net/http"

	// "github.com/xnpltn/instagram-backend/internal/handlers"
	"github.com/xnpltn/instagram-backend/internal/models"
	"github.com/xnpltn/instagram-backend/internal/utils"
)

type authHandler func(http.ResponseWriter, *http.Request, models.DBUser)

func AuthMiddleware(handler authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()
		// tokenKey := cookies[0].Name

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
		// apiKey, err := utils.GetAPiKey(r.Header)
		// if err != nil{
		// 	utils.RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		// 	return
		// }
		
		user, err := utils.VerifyJWT(tokenKey)
		if err != nil{
			utils.RespondWithError(w, 400, fmt.Sprintf("Cuuld not get user: %v", err))
			return
		}
		
		handler(w, r, user)
	}
}