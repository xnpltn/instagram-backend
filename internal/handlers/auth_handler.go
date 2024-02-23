package handlers

import (
	
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"database/sql"

	"github.com/xnpltn/instagram-backend/internal/database"
	"github.com/xnpltn/instagram-backend/internal/models"
	"github.com/xnpltn/instagram-backend/internal/utils"
)

func AuthHandler(w http.ResponseWriter, _ *http.Request){
	w.Write([]byte("Authentication"))
}

type User struct{}

func NewUser() *User{
	return &User{}
}


func(u *User) HandlerCreateUser(w http.ResponseWriter, r *http.Request){
	insertStmt := `
		INSERT INTO users (name, username, password)
		VALUES ($1, $2, $3)
		RETURNING id, name, username, password
	`
	ceStmt := `
		SELECT username FROM users WHERE username=$1
	`
	dbUser := models.DBUser{}
	decoder := json.NewDecoder(r.Body)
	params := models.CreateUserParams{}
	error := decoder.Decode(&params)
	if error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error Persing JSON: %v", error))
		return
	}
	db, err := database.Connect()
	errDB := db.QueryRow(ceStmt, params.Usename).Scan(&dbUser.Usename,)
	if errDB != nil{
		if !errors.Is(err, sql.ErrNoRows){
			log.Print("error: ", errDB)
		}
	}
	
	if err != nil {
		log.Fatal("Error initiating database", err)
	}
	params.Password, err = utils.HashPassword(params.Password)
	if err != nil{
		log.Fatal("error hashing password")
	}
	if dbUser.Usename == params.Usename{
		utils.RespondWithError(w, 400, "Username Taken")
	}else{
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
	
}

func(u *User) HandlerLoginUser(w http.ResponseWriter, r *http.Request,){
	
	getUserStmt := `
	SELECT id, username, name,  password FROM users WHERE username = $1;
	`
	user := models.DBUser{}
	decoder := json.NewDecoder(r.Body)
	params := models.LoginUserParams{}
	databaseUser := models.LoginUserParams{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error Persing JSON: %v", err))
		return
	}
	db, err := database.Connect()
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error Persing JSON: %v", err))
		return
	}
	errDB := db.QueryRow(getUserStmt, params.Usename).Scan(&databaseUser.ID,&databaseUser.Usename, &user.Name, &databaseUser.Password)
	if errDB !=nil{
		if errDB.Error() == "sql: no rows in result set"{
			utils.RespondWithError(w, 404, "User Not Found")
		}
	}

	

	
	if !utils.CheckPasswordHash(params.Password, databaseUser.Password){
		utils.RespondWithError(w, 401, "Invalid username or password")
	}else{
		
		user.ID = databaseUser.ID
		user.Usename = databaseUser.Usename
		tokenSting, err := utils.GenerateJWT(user)
		if err != nil{
			log.Fatal("error occured", err)
		}

		cookie := &http.Cookie{
			Name:  "Token",
			Value: tokenSting,
			Path:  "/",
			MaxAge: 36000, 
			HttpOnly: true,
			Secure: true, 
			SameSite: http.SameSiteLaxMode,
		}
		
		http.SetCookie(w, cookie)
		utils.RespondWithJson(w, http.StatusOK, "login successfully")
	}	
}