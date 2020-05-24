package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/connerj70/seva/internal/app/seva"
	"github.com/connerj70/seva/internal/mongo/user"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	DB           *mongo.Database
	JWTSecretKey string
}

// Retrieve will retrieve a user
func (u *User) Retrieve(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := params.ByName("id")

	user, err := user.Retrieve(u.DB, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get user %s", err), 500)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "failed to marshal user", 500)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
	w.Write(userJSON)
	return
}

// Create will create a new user
func (u *User) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var newUser seva.NewUser

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "failed to decode body", 500)
		return
	}

	// Call user create
	userID, err := user.Create(u.DB, newUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create user %s", err), 500)
		return
	}

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"id": %q}`, userID)))

	return
}

func (u *User) LogIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var getUser seva.User
	err := json.NewDecoder(r.Body).Decode(&getUser)
	if err != nil {
		http.Error(w, "failed to decode user", 500)
		return
	}

	userWithJWT, err := user.LogIn(u.DB, getUser.Email, getUser.Password, u.JWTSecretKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to log user in %s", err), 500)
		return
	}

	userWithJWTJSON, err := json.Marshal(userWithJWT)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal return user %s", err), 500)
		return
	}

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	w.Write(userWithJWTJSON)
	return
}
