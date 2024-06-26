package handlers

import (
	"github.com/charmbracelet/log"
	"net/http"
)

func (_ *Env) Ping(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Error(err)
	}
}
