package server

import (
	"context"
	"customer-board/db"
	"customer-board/router"
	"errors"
	"github.com/charmbracelet/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	db         *db.DB
	httpServer *http.Server
}

func New() *Server {
	d, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	httpServer := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router.New(),
	}

	return &Server{d, &httpServer}
}

func (s *Server) Start() error {
	go func() {
		log.Infof("Starting server on port %s", os.Getenv("PORT"))
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info("Shutting down server...")

	s.db.Close()
	return s.httpServer.Shutdown(ctx)
}
