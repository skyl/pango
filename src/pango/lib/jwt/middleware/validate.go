package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	// TODO: https://github.com/dgrijalva/jwt-go/blob/master/MIGRATION_GUIDE.md
	// "github.com/dgrijalva/jwt-go/request"

	"pango/lib/jwt/keys"
)

func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return keys.GetVerifyKey(), nil
	})
	fmt.Println(err)
	if err == nil {
		if token.Valid{
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorised access to this resource")
	}
}
