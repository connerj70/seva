package routes

import (
	"net/http"

	"github.com/connerj70/seva/internal/app/wire"
)

func RegisterSandwich() {
	sandwich := wire.WireSandwich()
	http.HandlerFunc("/sandwich", sandwich.Post)
}
