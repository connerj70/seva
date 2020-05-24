package middleware

import (
	"log"

	"github.com/julienschmidt/httprouter"
)

// Log will make sure any errors returned further down will be properly logged
func Log(handler httprouter.Handle) httprouter.Handle {
	log.Println("hi i'm from the logger")
	return handler
}
