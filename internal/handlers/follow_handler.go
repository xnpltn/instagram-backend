package handlers

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/xnpltn/instagram-backend/internal/database"
	"github.com/xnpltn/instagram-backend/internal/models"
	"github.com/xnpltn/instagram-backend/internal/utils"
)

type Follow struct{}

func NewFollow() *Follow {
	return &Follow{}
}

func (p *Follow) FollowUser(w http.ResponseWriter, r *http.Request, user models.DBUser){
	db, err := database.Connect()

	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	
	vars := mux.Vars(r)

	id := vars["userID"]

	if id == user.ID.String(){
		utils.RespondWithError(w, 400, "Can't follow yourself")
	}else{

		follow, err :=utils.FollowUserByID(db, user, id)
		if err != nil{
			log.Fatal("error ", err)
		}
	
		utils.RespondWithJson(w, 201, follow)
	}
}


func(p *Follow) UnfollowUser(w http.ResponseWriter, r *http.Request, user models.DBUser){
	db, err := database.Connect()

	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	
	vars := mux.Vars(r)

	id := vars["userID"]

	if id == user.ID.String(){
		utils.RespondWithError(w, 400, "Can't unfollow yourself")
	}else{
		err :=utils.UnfollowUserByID(db, user, id)
		if err != nil{
			utils.RespondWithError(w, 400, "error getting a user")
		}else{
			utils.RespondWithJson(w, 201, "UnFOllowed")
		}
	}
}

func (p *Follow)GetAllFollowing(w http.ResponseWriter, r *http.Request, user models.DBUser){
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	follwong, err := utils.GetFollowing(db, user)
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	utils.RespondWithJson(w, 200, follwong)
}