package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vvvkkkggg/xp-homework-service/internal/application"
)

type Handler struct {
	app *application.Application
}

func NewHandler(app *application.Application) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) home(ctx *gin.Context) {
	/*
		name := c.Param("name")
				c.String(http.StatusOK, "Hello %s", name)
	*/

	userID := ctx.Param("user")

	resp, err := h.app.Home(ctx.Request.Context(), userID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) task(ctx *gin.Context) {
	userID := ctx.Param("user")
	taskID := ctx.Param("task")

	resp, err := h.app.Task(ctx.Request.Context(), req)
	if err != nil {
		log.Error("unable to process request")
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

//func (h *Handler) submit(ctx *gin.Context) {
//
//	userID := ctx.Param("user")
//	taskID := ctx.Param("task")
//	var req model.ImagineRequest
//
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.Error(err)
//		ctx.Status(http.StatusBadRequest)
//		return
//	}
//
//	if err != nil {
//		h.logger.Error("unable to process request", zap.Error(err))
//		ctx.Error(err)
//		ctx.Status(http.StatusInternalServerError)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, resp)
//}
