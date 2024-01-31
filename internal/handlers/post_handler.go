package handlers

import (
	"log"
	"net/http"
	"github.com/xnpltn/codegram/internal/models"
	"github.com/xnpltn/codegram/internal/utils"
	"github.com/gorilla/mux"
	"github.com/xnpltn/codegram/internal/database"
)

type Post struct{}

func NewPost() *Post{
	return &Post{}
}


func (p *Post) CreatePosts(w http.ResponseWriter, r *http.Request, user models.DBUser){
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	post, err := utils.CreatePost(r, db, user)
	if err != nil {
		utils.RespondWithError(w, 404, err.Error())
	}
	utils.RespondWithJson(w, 201, post)
	
}

func (p *Post) GetPosts(w http.ResponseWriter, _ *http.Request){
	db, err := database.Connect()
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

func (p *Post) GetPostByID(w http.ResponseWriter, r *http.Request){
	db, err := database.Connect()
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

func (p *Post) DeletePostByID(w http.ResponseWriter, r *http.Request, user models.DBUser){
	db, err := database.Connect()
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

func (p *Post) EditPostByID(w http.ResponseWriter, r *http.Request, user models.DBUser){
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	vars := mux.Vars(r)
	id := vars["postID"]
	post, err := utils.EditPostByID(db, id, &user, r)
	if err != nil {
		utils.RespondWithError(w, 404, "Can't Update Post")
	}

	if post[0].ImageURL == "" && post[0].Description == ""{
		utils.RespondWithError(w, 400, "Edit your posts")
	}else{
		utils.RespondWithJson(w, 201, post)
	}
	
	
}
