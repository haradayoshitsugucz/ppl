package router

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/haradayoshitsugucz/purple-server/api/di"
	"github.com/haradayoshitsugucz/purple-server/api/presentation/middleware"
	"github.com/haradayoshitsugucz/purple-server/config"
	"github.com/haradayoshitsugucz/purple-server/logger"
)

var (
	cli = &http.Client{Timeout: 10 * time.Second}
)

// Run 起動
func Run(conf config.Config, logFileName string) error {

	logger.InitLogger(conf, logFileName)
	logger.GetLogger().Info(fmt.Sprintf("Purple API is Running at %v Port ...", conf.AppSetting().Port))
	logger.GetLogger().Info(conf.Description())

	return router(conf)
}

// router ルーター
func router(conf config.Config) error {

	handler := di.InitHandler(conf, time.Now, cli)
	ctrl := handler.Ctrl
	middle := handler.Middle

	appRouter := chi.NewRouter()
	middleware.Use(appRouter, conf)

	contextPath := fmt.Sprintf("/%v", conf.AppSetting().ContextPath)
	appRouter.Get(contextPath, func(aWriter http.ResponseWriter, aRequest *http.Request) {
		http.Redirect(aWriter, aRequest, contextPath+"/", http.StatusMovedPermanently)
	})

	appRouter.Mount(contextPath+"/", func() http.Handler {
		mainRouter := chi.NewRouter()
		mainRouter.Mount("/v1", func() http.Handler {
			r := chi.NewRouter()
			r.Use(middle.TimeMiddle.Now())
			r.Get("/status", ctrl.StatusController.Ping())
			r.Get("/products/{product_id}", ctrl.ProductController.GetProduct(conf))
			r.Delete("/products/{product_id}", ctrl.ProductController.DeleteProduct(conf))
			r.Put("/products/{product_id}", ctrl.ProductController.EditProduct(conf))
			r.Post("/products", ctrl.ProductController.AddProduct(conf))
			r.Get("/products", ctrl.ProductController.GetProductsByName(conf))
			return r
		}())
		return mainRouter
	}())

	// デバッグ用
	printRouter(appRouter)

	return http.ListenAndServe(fmt.Sprintf(":%d", conf.AppSetting().Port), appRouter)
}

// printRouter デバッグ用
func printRouter(r *chi.Mux) {

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		logger.GetLogger().Info(fmt.Sprintf("%s %s", method, route))
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		logger.GetLogger().Error(fmt.Sprintf("Logging err: %s", err.Error()))
	}
}
