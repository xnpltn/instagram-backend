package handlers

import (
	"log"
	"net/http"
	"github.com/xnpltn/codegram/internal/models"
	"github.com/xnpltn/codegram/internal/utils"
	"github.com/gorilla/mux"
)

func CreatePosts(w http.ResponseWriter, r *http.Request, user models.DBUser){
	db, err := models.InitDB()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	post, err := utils.CreatePost(w, r, db, user)
	if err != nil {
		utils.RespondWithError(w, 404, err.Error())
	}
	utils.RespondWithJson(w, 201, post)
	
}

func GetPosts(w http.ResponseWriter, _ *http.Request){
	db, err := models.InitDB()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	posts, err := utils.GetAllPosts(db)

	if err != nil {
		log.Fatal("Getting all posts", err)
	}
	utils.RespondWithJson(w, 200, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request, ){
	db, err := models.InitDB()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	id := vars["postID"]
	post, err := utils.GetPost(db, id)
	if err != nil {
		log.Fatal("error", err)
	}
	utils.RespondWithJson(w, 200, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request, user models.DBUser){
	db, err := models.InitDB()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	id := vars["postID"]
	post, err := utils.DeletePost(db, id, &user)
	if err != nil {
		log.Fatal("error", err)
	}
	utils.RespondWithJson(w, 204, post)
}
