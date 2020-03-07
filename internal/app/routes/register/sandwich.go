package register

import (
	"net/http"

	"github.com/connerj70/seva/internal/app/middleware"
)

func RegisterSandwich() {
	sandwich := WireSandwich()
	mid := &middleware.AuthMiddleware{Next: sandwich.Post}
	http.Handle("/sandwich", mid)
}
