package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/haradayoshitsugucz/purple-server/config"
	"github.com/haradayoshitsugucz/purple-server/logger"
	"moul.io/chizap"
)

// Use chi側のmiddleware
func Use(app *chi.Mux, conf config.Config) http.Handler {
	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	if conf.LoggerSetting().RequestOutput {
		app.Use(chizap.New(logger.GetLogger(), &chizap.Opts{
			WithReferer:   true,
			WithUserAgent: true,
		}))
	}
	app.Use(GlobalRecoverer)
	app.Use(middleware.NoCache)
	return app
}
