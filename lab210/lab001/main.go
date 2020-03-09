package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"log"
	"net/http"
	"time"
)

const (
	SecretKey = "lab001"
)

func main() {
	startServer()
}

func startServer() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/resource", resourceHandler)

	log.Println("Now listening...")
	http.ListenAndServe(":8080", nil)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "no access to resource")
	} else {
		if token.Valid {
			fmt.Fprint(w, "token is valid")
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	//sign
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error sign token")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tokenString)
}
