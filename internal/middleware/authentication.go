package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// Authenticate will parse and validate the jwt before calling the next handler
func Authenticate(handler httprouter.Handle, jwtSecretKey string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		jwtString := r.Header.Get("Authorization")
		if jwtString == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Parse the jwt
		_, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(jwtSecretKey), nil
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("jwt issue: %s", err), 500)
			return
		}
		// If the user is signed in, call the next handler
		handler(w, r, ps)
	}
}
