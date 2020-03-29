package register

import (
	"net/http"

	"github.com/connerj70/seva/internal/app/middleware"
)

func RegisterSandwich() {
	sandwich := WireSandwich()
	authMid := &middleware.AuthMiddleware{Next: sandwich.Post}
	cTypeMid := &middleware.ContentTypeMiddleware{authMid.ServeHTTP}
	http.Handle("/sandwich", cTypeMid)
}
