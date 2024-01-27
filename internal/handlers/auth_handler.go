package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/xnpltn/codegram/internal/models"
	"github.com/xnpltn/codegram/internal/utils"
)

func AuthHandler(w http.ResponseWriter, _ *http.Request){
	w.Write([]byte("Authentication"))
}




func HandlerCreateUser(w http.ResponseWriter, r *http.Request){
	insertStmt := `
		INSERT INTO users (name, username, password)
		VALUES ($1, $2, $3)
		RETURNING id, name, username, password
	`
	
	decoder := json.NewDecoder(r.Body)
	params := models.CreateUserParams{}
	error := decoder.Decode(&params)
	if error != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error Persing JSON: %v", error))
		return
	}
	db, err := models.InitDB()
	
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	err = db.QueryRow(insertStmt, params.Name, params.Usename, params.Password).Scan(&params.ID, &params.Name, &params.Usename, &params.Password)
	if err != nil{
		log.Fatal("Errpr occured adding user to the database", err)
	}
	user := models.DBUser{
		ID: params.ID,
		Name: params.Name,
		Usename: params.Usename,
	}
	utils.RespondWithJson(w, 201, user)
	db.Close()
}

func HandlerLoginUser(w http.ResponseWriter, r *http.Request){
	
	getUserStmt := `
	SELECT username, password FROM users WHERE username = $1;
	`

	decoder := json.NewDecoder(r.Body)
	params := models.LoginUserParams{}
	databaseUser := models.LoginUserParams{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error Persing JSON: %v", err))
		return
	}
	db, err := models.InitDB()
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error Persing JSON: %v", err))
		return
	}
	err = db.QueryRow(getUserStmt, params.Usename).Scan(&databaseUser.Usename, &databaseUser.Password)
	if err !=nil{
		log.Fatal("error signing in", err)
	}

	if databaseUser.Password != params.Password{
		utils.RespondWithError(w, 404, "Invalid username or password")
	}
	
	fmt.Println("database user",databaseUser)
}