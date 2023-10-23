package rest

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/vvvkkkggg/xp-homework-service/internal/application"
)

func CreateAndServe(ctx context.Context, app *application.Application, router *Router) error {
	server := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", "80"),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("unable to serve http")
	}

	return nil
}
