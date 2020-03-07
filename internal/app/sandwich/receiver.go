package sandwich

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BusinessAdapter interface {
	Post(*Sandwich) error
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
