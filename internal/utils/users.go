package utils

import (
	"database/sql"
	"github.com/xnpltn/instagram-backend/internal/models"
	"log"
)

func GetUserByID(db *sql.DB, id string) ([]models.DBUser, error) {
	stmt, err := db.Prepare("SELECT id, name, username FROM users WHERE id=$1")
	if err != nil {
		log.Fatal("error 1 ", err)
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal("error 2 ", err)
	}

	user, err := getUsersFromDB(rows)
	if err != nil {
		log.Fatal("error 3 ", err)
	}
	return user, nil
}

func getUsersFromDB(rows *sql.Rows) ([]models.DBUser, error) {
	items := []models.DBUser{}
	for rows.Next() {
		var user models.DBUser
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Usename,
		); err != nil {
			log.Fatal("error 4 ", err)
		}
		items = append(items, user)
	}

	return items, nil
}
