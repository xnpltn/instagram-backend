package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/xnpltn/codegram/internal/models"
)

func GenerateJWT(username models.DBUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": username.ID,
		"name": username.Name,
		"username": username.Usename,
		"iat":time.Now().Unix(),
	})
	secret := []byte("secret")

	tokenSting, err := token.SignedString(secret)

	return tokenSting, err
}


func VerifyJWT( token string) (models.DBUser, error){
	secret := []byte("secret")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil{
		log.Fatal("erro persing jwt")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok{
		fmt.Println("Error")
	}

	idString, _ := claims["id"].(string)

	idUUID, err := uuid.Parse(idString)
	if err != nil {
		log.Fatal("eror", err)
	}
	user := models.DBUser{
		ID: idUUID,
		Name: claims["name"].(string),
		Usename: claims["username"].(string),
	}

	return user, nil
}


func GetAPiKey(header http.Header) (string, error){
	val := header.Get("Authorization")
	if val == ""{
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "Token" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil

}

