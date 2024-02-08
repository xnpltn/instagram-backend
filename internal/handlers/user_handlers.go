package handlers

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/xnpltn/instagram-backend/internal/database"
	"github.com/xnpltn/instagram-backend/internal/utils"
)



func (p *User) GetUserByID(w http.ResponseWriter, r *http.Request){
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	id := vars["userID"]
	user, err := utils.GetUserByID(db, id)
	if err != nil {
		utils.RespondWithError(w, 404, err.Error())
	}
	utils.RespondWithJson(w, 201, user)
	
}