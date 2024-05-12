package main

import (
	"customer-board/db"
	"customer-board/server"
	"github.com/charmbracelet/log"
)

type App struct {
	db *db.DB
}

func main() {
	s := server.New()

	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
