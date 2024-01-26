package models

import(
	"log"
	"database/sql"
)

func InitDB() (*sql.DB, error){
	db, err := sql.Open("postgres", "postgresql://postgres:qwerty@localhost:5432/codegram?sslmode=disable")
	if err != nil{
		log.Fatal("Error Occured Connectiong to database", err)
	}
	stmt, err := db.Prepare(
		`
		CREATE TABLE IF NOT EXISTS users (
			id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
			name VARCHAR(255),
			username VARCHAR(255) UNIQUE,
			password VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		
		`,
	)
	if err != nil{
		log.Fatal("Error occured creating Tables", err)
	}

	res, err:= stmt.Exec()
	if err != nil{
		log.Fatal(err) 
	}
	log.Println(res)
	log.Println("Database initiated succefull")
	return db, nil
}