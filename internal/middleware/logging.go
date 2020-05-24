package middleware

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type statusRecorder struct {
	http.ResponseWriter
	status  int
	message string
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *statusRecorder) Write(b []byte) (int, error) {
	rec.message = string(b)
	return rec.ResponseWriter.Write(b)
}

// Log will make sure any errors returned further down will be properly logged
func Log(log *log.Logger, handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		rec := statusRecorder{w, 200, ""}
		handler(&rec, r, p)

		if rec.status != 200 {
			log.Println(rec.status, rec.message)
		}
	}
}
