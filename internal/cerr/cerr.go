package cerr

import (
	"encoding/json"
	"net/http"
)

// InternalError will store information about errors that happen in our application.
type InternalError struct {
	ID         string `json:",omitempty"`
	statusCode int
	Header     string
	Detail     string
	Err        error
}

func NewInternalError(statusCode int, header, detail string, err error) InternalError {
	return InternalError{
		statusCode: statusCode,
		Header:     header,
		Detail:     detail,
		Err:        err,
	}
}

func (ie InternalError) Error() string {
	ieBytes, _ := json.Marshal(ie)
	return string(ieBytes)
}

// Unwrap will return the inner error
func (ie InternalError) Unwrap() error {
	return ie.Err
}

// SetStatusCode will be used to set the status code of the InternalError struct.
func (ie *InternalError) SetStatusCode(statusCode int) {
	ie.statusCode = statusCode
}

// SetWriterStatusCode will call the response writers WriteHeader method with the InternalError's status code
// if it is not zero, otherwise it will write a status code of 500 internal server error.
func (ie *InternalError) SetWriterStatusCode(w http.ResponseWriter) {
	if ie.statusCode != 0 {
		w.WriteHeader(ie.statusCode)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
