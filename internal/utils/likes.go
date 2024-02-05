package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/xnpltn/codegram/internal/models"
)

func LikePost(db *sql.DB, user models.DBUser, id string) (models.DBLike, error){
	sqlStatement := `
		INSERT INTO likes (post_id, user_id) 
		VALUES ($1, $2)
		RETURNING post_id, user_id;
	`
	like := models.DBLike{}

	post, err := GetPost(db, id)

	if err != nil{
		log.Fatalln("error, err")
	}
	if len(post) < 1{
		return like, fmt.Errorf("post with id %s does not exist", id)
	}else{
		err = db.QueryRow(sqlStatement, id, user.ID  ).Scan(&like.PostID, &like.UserID)
		if err != nil{
			log.Println("error1likes: ", err)
		}
		return like, nil
	}
}


func UnlikePost(db *sql.DB, user models.DBUser, id string)(models.DBLike, error){
	sqlStmt := `
		DELETE FROM likes WHERE post_id = $1 AND user_id = $2
		RETURNING post_id, user_id;
	`
	like := models.DBLike{}
	err := db.QueryRow(sqlStmt, id, user.ID).Scan(&like.PostID, &like.UserID)
	if err != nil{
		log.Println("error Unlike: ", err)
	}
	return like, nil

}

func GetLikesCount(db *sql.DB, id string)( int64, error){
	sqlStmt := `
		SELECT COUNT(*) FROM likes WHERE post_id = $1;
	`
	var likes int64
	err := db.QueryRow(sqlStmt, id).Scan(&likes)

	if err!= nil{
		log.Fatalln("error count", err)
	}

	return likes, nil
}