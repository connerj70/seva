package routes

import (
	"net/http"

	"github.com/connerj70/seva/internal/app"
)

// RegisterNoAuth will initialize the no auth routes
func RegisterNoAuth() {
	noauth := app.WireNoAuth()
	http.HandleFunc("/noauth/register", noauth.Register)
}
