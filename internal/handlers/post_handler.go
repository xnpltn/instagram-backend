package handlers

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/xnpltn/codegram/internal/models"
	"github.com/xnpltn/codegram/internal/utils"
)

func CreatePosts(w http.ResponseWriter, r *http.Request, user models.DBUser){
	sqlStatement := `
		INSERT INTO posts (user_id, image_url, description)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, image_url, description;
	`
	db, err := models.InitDB()
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	decoder := json.NewDecoder(r.Body)
	params := models.CreatePostParams{}
	post := models.DBPost{}
	error := decoder.Decode(&params)
	if error != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error Persing JSON: %v", error))
		return
	}
	
	err = db.QueryRow(sqlStatement, user.ID, params.ImageURL, params.Description).Scan(&post.ID, &post.UserID, &post.ImageURL, &post.Description)
	
	if err != nil{
		log.Fatal("error", err)
	}
	utils.RespondWithJson(w, 201, post)
	
}

func GetPosts(w http.ResponseWriter, _ *http.Request, _ models.DBUser){
	w.Write([]byte("get posts"))
}
