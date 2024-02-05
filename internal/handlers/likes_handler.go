package handlers

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/xnpltn/codegram/internal/database"
	"github.com/xnpltn/codegram/internal/models"
	"github.com/xnpltn/codegram/internal/utils"
)

type Like struct {}


func NewLike() *Like{
	return &Like{}
}

func (l *Like ) LikePostByID(w http.ResponseWriter, r * http.Request, user models.DBUser){
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	id := vars["postID"]
	like, err :=utils.LikePost(db, user, id)
	if err != nil {
		utils.RespondWithError(w, 404, err.Error())
	}else{
		utils.RespondWithJson(w, 201, like)
	}
}

func (l *Like ) UnikePostByID(w http.ResponseWriter, r * http.Request, user models.DBUser){
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	id := vars["postID"]
	like, err :=utils.UnlikePost(db, user, id)
	if err != nil {
		utils.RespondWithError(w, 404, err.Error())
	}else{
		utils.RespondWithJson(w, 201, like)
	}
}


func (l *Like ) GetLikesCountByID(w http.ResponseWriter, r * http.Request){
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	id := vars["postID"]

	likes, err := utils.GetLikesCount(db, id)
	if err != nil {
		log.Fatal("Error retunning likes", err)
	}

	utils.RespondWithJson(w, 200, likes)
}