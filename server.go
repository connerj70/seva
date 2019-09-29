package main

import (
	"fmt"
	"io"
	"net/http"

"github.com/connerj70/seva/noauth"
)

func main() {
	fmt.Println("hi")

	http.HandleFunc("/", a)

	http.ListenAndServe(":8080", nil)
}

func a(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "you hit a!!!")
}
