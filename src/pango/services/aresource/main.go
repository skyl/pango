package main

import (
  "log"
  "net/http"

  "github.com/codegangsta/negroni"

  "pango/lib/jwt/middleware"
  "pango/lib/web/response"
)

type Response struct {
	Data	string	`json:"data"`
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{"Gained access to protected resource"}
	response.Json(res, w)
}

func main() {
  // PROTECTED ENDPOINTS
  http.Handle("/resource/", negroni.New(
    negroni.HandlerFunc(middleware.ValidateToken),
    negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
  ))
  log.Println("Now listening on 8081...")
  // TODO: how to do routing? :/
  http.ListenAndServe(":8081", nil)
}
