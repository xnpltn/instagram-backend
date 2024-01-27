package handlers

import (
	"log"
	// "fmt"
	"net/http"
	// "encoding/json"
	"github.com/xnpltn/codegram/internal/models"
	"github.com/xnpltn/codegram/internal/utils"
	"github.com/gorilla/mux"
)

func CreatePosts(w http.ResponseWriter, r *http.Request, user models.DBUser){
	// sqlStatement := `
	// 	INSERT INTO posts (user_id, image_url, description)
	// 	VALUES ($1, $2, $3)
	// 	RETURNING id, user_id, image_url, description;
	// `
	db, err := models.InitDB()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	defer db.Close()
	
	// decoder := json.NewDecoder(r.Body)
	// params := models.CreatePostParams{}
	// post := models.DBPost{}
	// error := decoder.Decode(&params)
	// if error != nil {
	// 	utils.RespondWithError(w, 400, fmt.Sprintf("Error Persing JSON: %v", error))
	// 	return
	// }
	
	// err = db.QueryRow(sqlStatement, user.ID, params.ImageURL, params.Description).Scan(&post.ID, &post.UserID, &post.ImageURL, &post.Description)
	
	// if err != nil{
	// 	log.Fatal("error", err)
	// }
	post, err := utils.CreatePost(w, r, db, user)
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
