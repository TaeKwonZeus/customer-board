package main

import (
	"customer-board/src/router"
	"github.com/charmbracelet/log"
	"net/http"
	"os"
)

func main() {
	r := router.New()

	log.Infof("Starting server on port %s", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal(err)
	}
}
