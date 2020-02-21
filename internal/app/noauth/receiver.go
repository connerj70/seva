package noauth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BusinessAdapter will hold the method signature for talking with the business layer
type BusinessAdapter interface {
	Register(*User) error
}

// Receiver will hold the adapter for the business layer
type Receiver struct {
	Business BusinessAdapter
}

// Register will handle all user registration
func (rec *Receiver) Register(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	if err := rec.Business.Register(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
