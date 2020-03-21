package sandwich

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/connerj70/seva/internal/cerr"
)

type BusinessAdapter interface {
	Post(*Sandwich) error
	Put(*Sandwich) error
}

type Receiver struct {
	Business BusinessAdapter
}

func (r *Receiver) Post(w http.ResponseWriter, req *http.Request) {
	sandwich := &Sandwich{}
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(sandwich)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	sandwichBytes, err := json.Marshal(sandwich)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(sandwichBytes)
}

// Put will update a sandwich.
func (r *Receiver) Put(w http.ResponseWriter, req *http.Request) {
	sandwich := &Sandwich{}
	defer req.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(req.Body).Decode(sandwich)
	if err != nil {
		http.Error(w, fmt.Sprintf("there was a problem decoding the body %s", err), http.StatusInternalServerError)
		return
	}

	err = r.Business.Put(sandwich)
	if err != nil {
		var internalErr *cerr.InternalError
		if errors.As(err, &internalErr) {
			internalErr.SetWriterStatusCode(w)
			internalErrBytes, err := json.Marshal(internalErr)
			if err != nil {
				http.Error(w, fmt.Sprintf("unable to marshal internal error %s", err), http.StatusInternalServerError)
				return
			}
			w.Write(internalErrBytes)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
