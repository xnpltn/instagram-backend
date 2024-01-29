package database

import (
	"database/sql"
	"log"
)

func Connect() (*sql.DB, error){
	db, err := sql.Open("postgres", "postgresql://postgres:qwerty@localhost:5432/codegram?sslmode=disable")
	if err != nil{
		log.Fatal("Error Occured Connectiong to database", err)
	}
	_ , err = db.Exec(
		`
		CREATE TABLE IF NOT EXISTS users (
			id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL,
			username VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		
		CREATE TABLE IF NOT EXISTS posts (
            id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
            user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
            image_url VARCHAR(255) NOT NULL,
            description TEXT
        );
		`,
	)

	if err != nil{
		log.Fatal("Error occured creating Tables", err)
	}

	return db, nil
}