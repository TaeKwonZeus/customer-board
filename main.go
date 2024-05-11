package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"path/filepath"
)

const (
	distFolder = "web/dist"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("pong"))
		})
	})

	r.Get("/favicon.ico", iconHandler)
	r.Get("/assets/{asset}", assetHandler)
	r.Get("/*", frontendHandler)

	_ = http.ListenAndServe(":8080", r)
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.ReadFile(filepath.Join(distFolder, "favicon.ico"))
	_, _ = w.Write(f)
}

func assetHandler(w http.ResponseWriter, r *http.Request) {
	assetParam := chi.URLParam(r, "asset")

	http.ServeFile(w, r, filepath.Join(distFolder, "assets", assetParam))
}

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.ReadFile(filepath.Join(distFolder, "index.html"))
	_, _ = w.Write(f)
}
