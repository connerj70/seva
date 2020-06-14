package web

import "net/http"

// SuccessResponse will set a status code of 200, a content type of application/json and write the response bytes to the http.ResponseWriter
func SuccessResponse(w http.ResponseWriter, responseBytes []byte) {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	w.Write(responseBytes)
}
