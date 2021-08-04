package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/golang-jwt/jwt"
)

var mySigningKey = []byte("anything")

/* Homepage decorator to see if the token header is included in the request and also
   if the token is valid depending on the signingKey. */
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(
				r.Header["Token"][0],
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						//Errorf strings can't be capitalized.
						return nil, fmt.Errorf("error")
					}
					return mySigningKey, nil
				},
			)
			if err != nil {
				fmt.Fprint(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not authorized")
		}
	})
}

//Simple REST API.

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
	fmt.Println("Endpoint hit: homePage")
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	handleRequests()
}
