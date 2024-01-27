package models

import (
	"database/sql"
	"log"
)

func InitDB() (*sql.DB, error){
	db, err := sql.Open("postgres", "postgresql://postgres:qwerty@localhost:5432/codegram?sslmode=disable")
	if err != nil{
		log.Fatal("Error Occured Connectiong to database", err)
	}
	_ , err = db.Exec(
		`
		CREATE TABLE IF NOT EXISTS users (
			id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
			name VARCHAR(255),
			username VARCHAR(255) UNIQUE,
			password VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		
		CREATE TABLE IF NOT EXISTS posts (
            id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
            user_id UUID REFERENCES users(id),
            image_url VARCHAR(255),
            description TEXT
        );
		`,
	)

	if err != nil{
		log.Fatal("Error occured creating Tables", err)
	}

	return db, nil
}