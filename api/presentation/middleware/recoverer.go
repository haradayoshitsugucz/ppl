package middleware

// 下記ファイルのコピーをカスタマイズ
// The original work was derived from Goji's middleware, source:
// https://github.com/zenazn/goji/tree/master/web/middleware
// https://github.com/zenazn/goji/blob/master/web/middleware/recoverer.go

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/haradayoshitsugucz/purple-server/logger"
)

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
//
// Alternatively, look at https://github.com/pressly/lg middleware pkgs.
func GlobalRecoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {

				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					fmt.Fprintf(os.Stderr, "Panic: %+v\n", rvr)
					debug.PrintStack()
				}
				// logging
				logger.GetLogger().Error(fmt.Sprintf("Panic: %+v", rvr))
				logger.GetLogger().Error(fmt.Sprintf("Panic: %+v", string(debug.Stack())))

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
