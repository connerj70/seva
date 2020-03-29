package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type AuthMiddleware struct {
	Next http.HandlerFunc
}

type ContentTypeMiddleware struct {
	Next http.HandlerFunc
}

func (c *ContentTypeMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c.Next(w, r)
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if we have a jwt
	t := r.Header.Get("Authorization")
	if t == "" {
		w.WriteHeader(401)
		w.Write([]byte("missing authorization header, please log in"))
		return
	}

	// If we have a jwt let's check if it is valid
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "there was a problem parsing the authorization token %s", err)
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("It was valid!")
	} else {
		w.WriteHeader(401)
		w.Write([]byte("invalid authorization token"))
		return
	}

	// Call the next function
	m.Next(w, r)
}
