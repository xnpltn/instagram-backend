package utils

import (
	"database/sql"
	"log"
	"github.com/xnpltn/codegram/internal/models"
)

func FollowUserByID(db *sql.DB, user models.DBUser, id string) (models.DBFollow, error){
	sqlStatement := `
	INSERT INTO followers (follower_id, following_id) VALUES ($1, $2)
	RETURNING follower_id, following_id;
	`
	follow := models.DBFollow{}
	dbUser, err := GetUserByID(db, id)
	if err != nil{
		log.Fatal("erroe", err)
	}
	
	db.QueryRow(sqlStatement, user.ID, dbUser[0].ID ).Scan(&follow.FollowerID, &follow.FollowingID)
	if err != nil{
		log.Fatal("error ", err)
	}
	return follow, nil
}


func UnfollowUserByID(db *sql.DB, user models.DBUser, id string) error{
	sqlStmt := `
		DELETE FROM followers where follower_id=$1 AND  following_id=$2;
	`
	dbUser, err := GetUserByID(db, id)
	if err != nil{
		log.Fatalln("erroe", err)
		return err
	}
	db.QueryRow(sqlStmt, user.ID, dbUser[0].ID)
	return nil
}


func GetFollowing(db *sql.DB, user models.DBUser) ([]models.DBFollow, error){
	sqlStmt := `
		SELECT * FROM followers WHERE follower_id=$1
	`
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal("error", err)
	}
	rows, err := stmt.Query(user.ID)
	if err != nil {
		log.Fatal("error", err)
	}

	followings , err := getFollowerItemsFromDB(rows)
	if err != nil{
		log.Fatalln("error", err)
	}
	return followings, nil
}


func getFollowerItemsFromDB(rows *sql.Rows) ([]models.DBFollow, error){
	follows := []models.DBFollow{}
	for rows.Next(){
		var follow models.DBFollow
		if err := rows.Scan(
			&follow.FollowerID,
			&follow.FollowingID,
		); err != nil {
			log.Fatal("error", err)
		}
		follows = append(follows, follow)
	}

	return follows, nil
}