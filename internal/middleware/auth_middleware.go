package middleware

import (
	"fmt"
	"log"
	"net/http"
	"github.com/xnpltn/codegram/internal/models"
	"github.com/xnpltn/codegram/internal/utils"
)

type authHandler func(http.ResponseWriter, *http.Request, models.DBUser)

func AuthMiddleware(handler authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := utils.GetAPiKey(r.Header)
		if err != nil{
			utils.RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := utils.VerifyJWT(apiKey)
		if err != nil{
			log.Fatal("Error")
		}
		
	
		if err != nil{
			utils.RespondWithError(w, 400, fmt.Sprintf("Cuuld not get user: %v", err))
			return
		}
		
		handler(w, r, user)
	}
}