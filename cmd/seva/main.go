package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/connerj70/seva/internal/app/routes"
)

func main() {
	routes.RegisterRoutes()

	fmt.Println("Seva Is Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
