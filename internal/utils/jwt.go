package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/xnpltn/instagram-backend/internal/models"
)

func GenerateJWT(username models.DBUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": username.ID,
		"name": username.Name,
		"username": username.Usename,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	secret := []byte("3729c8d0d8682548d4c918d0655a950caeadad4d391b7f319d5a839231f92231")

	tokenSting, err := token.SignedString(secret)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(tokenSting))

	return encodedToken, err
}


func VerifyJWT( token string) (models.DBUser, error){
	secret := []byte("3729c8d0d8682548d4c918d0655a950caeadad4d391b7f319d5a839231f92231")
	decodedToken, err:= base64.StdEncoding.DecodeString(token)
	if err != nil{
		log.Println("error", err)
	}
	parsedToken, err := jwt.Parse(string(decodedToken), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil{
		log.Println("erro persing jwt:", err)
		return models.DBUser{}, err
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

