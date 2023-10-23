package rest

import "github.com/vvvkkkggg/xp-homework-service/internal/application"

type Handler struct {
	svc *application.Application
}

func NewHandler(svc *application.Application) *Handler {
	return &Handler{
		svc: svc,
	}
}
