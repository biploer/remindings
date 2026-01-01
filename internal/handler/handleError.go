package handler

import (
	"log"
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func handle(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handleError(w, r, err)
		}
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	_ = w
	_ = r
	log.Println(err)
}
