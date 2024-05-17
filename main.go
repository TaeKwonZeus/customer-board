package main

import (
	"context"
	"customer-board/db"
	"customer-board/handlers"
	"customer-board/router"
	"errors"
	"github.com/charmbracelet/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	d, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	env := handlers.Env{DB: d}

	httpServer := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router.New(env),
	}

	go func() {
		log.Infof("Starting server on port %s: http://localhost:%s", os.Getenv("PORT"), os.Getenv("PORT"))
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info("Shutting down server...")

	if err = httpServer.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server forced to shutdown:", err)
	}
}
