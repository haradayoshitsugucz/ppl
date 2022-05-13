package di

import (
	"github.com/haradayoshitsugucz/purple-server/api/presentation/controller"
	"github.com/haradayoshitsugucz/purple-server/api/presentation/middleware"
)

type ControllerAndMiddleware struct {
	Ctrl   *controller.Controller
	Middle *middleware.Middleware
}
