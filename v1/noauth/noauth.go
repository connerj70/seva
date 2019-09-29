package noauth

import (
	"io"
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "you hit the sign up method")
}
