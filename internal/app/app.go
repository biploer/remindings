package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Run(ctx context.Context) error {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, fellow!")
	})

	server := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  4 * time.Second,
		WriteTimeout: 4 * time.Second,
	}

	go func() {
		<-ctx.Done()
		log.Println("Server is shutting down...")
		server.Shutdown(ctx)
	}()

	log.Println("Server is starting...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}
