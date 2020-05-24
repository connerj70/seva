package handlers

import (
	"net/http"

	"github.com/connerj70/seva/internal/middleware"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
)

// Register registers all the handlers to the router
func Register(router *httprouter.Router, db *mongo.Database, jwtSecretKey string) {
	// Ping
	router.GET("/ping", ping)

	// User
	user := User{DB: db, JWTSecretKey: jwtSecretKey}
	router.GET("/user/:id", middleware.Authenticate(user.Retrieve, jwtSecretKey))
	router.POST("/user", user.Create)
	router.POST("/user/log_in", user.LogIn)
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("ping success!"))
}
