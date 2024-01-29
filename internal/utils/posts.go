package utils

import (
	"database/sql"
	"log"
	"github.com/xnpltn/codegram/internal/models"
	"net/http"
	"fmt"
	"os"
	"io"
)



func CreatePost(_ http.ResponseWriter, r *http.Request, db *sql.DB, user models.DBUser) (models.DBPost, error){
	r.ParseMultipartForm(20 * 1024 * 1024)
	sqlStatement := `
		INSERT INTO posts (user_id, image_url, description)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, image_url, description;
	`
	
	file, _, err := r.FormFile("image")

	if err != nil{
		fmt.Println("error occured", err)
	}

	imageDescription := r.FormValue("description")
	defer file.Close()
	if err = os.MkdirAll("uploads/"+ user.Usename, 0775); os.IsExist(err){
		fmt.Println("folder exists")
	}
	
	tempfile, err := os.CreateTemp("uploads/"+ user.Usename +"/", "uplaod-*.jpg")
	if err != nil{
		fmt.Println("error occured", err)
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil{
		fmt.Println("error occured", err)
	}
	if err != nil{
		fmt.Println("error occured", err)
	}
	tempfile.Write(fileBytes)
	post := models.DBPost{}
	imageUrl := tempfile.Name()

	err = db.QueryRow(sqlStatement, user.ID, imageUrl, imageDescription).Scan(&post.ID, &post.UserID, &post.ImageURL, &post.Description)
	
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

	items, err := getPostItemsFromDB(rows)

	if err != nil {
		log.Fatal("error", err)
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

	items, err := getPostItemsFromDB(rows)

	if err != nil {
		log.Fatal("error", err)
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
	items, err := getPostItemsFromDB(rows)

	if err != nil {
		log.Fatal("error", err)
	}

	return items, nil
}

func getPostItemsFromDB(rows *sql.Rows) ([]models.DBPost, error){
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