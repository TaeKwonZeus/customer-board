package router

import (
	"customer-board/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"path/filepath"
)

func New(env handlers.Env) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/api", apiRouter(env))
	r.Mount("/", frontendRouter())

	return r
}

func apiRouter(env handlers.Env) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)

	r.Get("/ping", env.Ping)

	return r
}

func frontendRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/favicon.ico", iconHandler)
	r.Get("/assets/{asset}", assetHandler)
	r.Get("/*", frontendHandler)

	return r
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, getFrontendAssets(), "favicon.ico")
}

func assetHandler(w http.ResponseWriter, r *http.Request) {
	assetParam := chi.URLParam(r, "asset")
	http.ServeFileFS(w, r, getFrontendAssets(), filepath.Join("assets", assetParam))
}

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, getFrontendAssets(), "index.html")
}
