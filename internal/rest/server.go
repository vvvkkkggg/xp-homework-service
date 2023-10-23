package rest

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
)

func CreateAndServe(ctx context.Context, router *Router) error {
	server := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", "80"),
		Handler: router,
	}

	log.Println("server is about to start")

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("unable to serve http")
	}

	return nil
}
