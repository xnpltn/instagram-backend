package utils

import (
	"database/sql"
	"log"
	"github.com/xnpltn/codegram/internal/models"
	"encoding/json"
	"net/http"
	"fmt"
)



func CreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB, user models.DBUser) (models.DBPost, error){
	sqlStatement := `
		INSERT INTO posts (user_id, image_url, description)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, image_url, description;
	`
	decoder := json.NewDecoder(r.Body)
	params := models.CreatePostParams{}
	post := models.DBPost{}
	error := decoder.Decode(&params)
	if error != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error Persing JSON: %v", error))
	}
	
	err := db.QueryRow(sqlStatement, user.ID, params.ImageURL, params.Description).Scan(&post.ID, &post.UserID, &post.ImageURL, &post.Description)
	
	if err != nil{
		log.Fatal("error", err)
	}
	return post, nil
}

func GetAllPosts(db *sql.DB) ([]models.DBPost, error){
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}

	items := []models.DBPost{}
	
	for rows.Next(){
		var post models.DBPost
		if err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.ImageURL,
			&post.Description,
		); err != nil {
			log.Fatal("error", err)
		}
		items = append(items, post)
	}

	return items, nil
}

func GetPost(db *sql.DB, id string) ([]models.DBPost, error){
	stmt, err := db.Prepare("SELECT * FROM posts WHERE id=$1")
	if err != nil {
		log.Fatal("error", err)
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal("error", err)
	}

	items := []models.DBPost{}
	for rows.Next(){
		var post models.DBPost
		if err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.ImageURL,
			&post.Description,
		); err != nil {
			log.Fatal("error", err)
		}
		items = append(items, post)
	}
	return items, nil
}


func DeletePost(db *sql.DB, id string, user *models.DBUser) ([]models.DBPost, error){
	stmt, err := db.Prepare("DELETE FROM posts WHERE id=$1 AND user_id=$2 RETURNING *;")
	if err != nil {
		log.Fatal("error", err)
	}
	rows, err := stmt.Query(id, user.ID)
	if err != nil {
		log.Fatal("error", err)
	}

	items := []models.DBPost{}
	for rows.Next(){
		var post models.DBPost
		if err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.ImageURL,
			&post.Description,
		); err != nil {
			log.Fatal("error", err)
		}
		items = append(items, post)
	}
	return items, nil
}