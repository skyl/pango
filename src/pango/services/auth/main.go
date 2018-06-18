package main

import (
	"log"
	"fmt"
	"time"
	"errors"
	"strings"
	"net/http"
	"encoding/json"

	"github.com/dgrijalva/jwt-go"

	"pango/lib/web/response"
	"pango/lib/jwt/keys"
)

type UserCredentials struct {
	Username	string  `json:"username"`
	Password	string	`json:"password"`
}

type Token struct {
	Token 	string    `json:"token"`
}

func validateCredentials(user UserCredentials) (err error){
	// TODO: use a database and hash and so forth
	if strings.ToLower(user.Username) != "skylar" {
		return errors.New("No such user")
	}
	if user.Password != "password1" {
		return errors.New("Wrong password")
	}
	return nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var usercred UserCredentials
	// decode request into UserCredentials struct
	err := json.NewDecoder(r.Body).Decode(&usercred)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request")
		return
	}
	err = validateCredentials(usercred)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Invalid Credentials!")
		return
	}
	// TODO: https://github.com/dgrijalva/jwt-go/blob/master/MIGRATION_GUIDE.md
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["iss"] = "admin"
	token.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	token.Claims["CustomUserInfo"] = struct {
		Name	string
		Role	string
	}{usercred.Username, "Member"}
	tokenString, err := token.SignedString(keys.GetSignKey())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
	}
	// create a token instance using the token string
	resp := Token{tokenString}
	response.Json(resp, w)
}

func main() {
	http.HandleFunc("/login", LoginHandler)
	log.Println("Now listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
