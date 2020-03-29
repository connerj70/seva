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
		return
	}

	err = r.Business.Post(sandwich)
	if err != nil {
		cErr := cerr.InternalError{}
		if errors.As(err, &cErr) {
			cErr.SetWriterStatusCode(w)
			cErr.Detail = fmt.Sprintf("failed business post: %s", cErr.Detail)
			w.Write([]byte(cErr.Error()))
			return
		}
		fmt.Fprintf(w, fmt.Sprintf("business failed to post sandwich %s", err))
		return
	}

	sandwichBytes, err := json.Marshal(sandwich)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

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
