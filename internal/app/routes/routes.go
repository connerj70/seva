package routes

import (
	"github.com/connerj70/seva/internal/app/routes/register"
)

//RegisterRoutes will set up all the handle funcs
func RegisterRoutes() {
	register.RegisterNoAuth()
	register.RegisterSandwich()
	register.RegisterTaco()
}
