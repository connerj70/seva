package register

import (
	"net/http"
)

// RegisterNoAuth will initialize the no auth routes
func RegisterNoAuth() {
	noauth := WireNoAuth()
	http.HandleFunc("/noauth/register", noauth.Register)
	http.HandleFunc("/noauth/log_in", noauth.LogIn)
}
