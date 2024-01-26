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
	`
	type parameters struct{
		Name string `json:"name"`
		Usename string `json:"username"`
		Password string `json:"password"`

	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	error := decoder.Decode(&params)
	if error != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error Persing JSON: %v", error))
		return
	}
	db, err := models.InitDB()
	
	if err != nil {
		log.Fatal("Error initializing tadabase", err)
	}
	_, errr := db.Exec(insertStmt, params.Name, params.Usename, params.Password)
	if errr != nil{
		log.Fatal("errpr occured adding user to the databse", errr)
	}
	utils.RespondWithJson(w, 201, params)
	db.Close()
}

func HandlerLoginUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Usename string `json:"username"`
		Password string `json:"password"`

	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	error := decoder.Decode(&params)
	if error != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error Persing JSON: %v", error))
		return
	}

	fmt.Println(params)
}