package rest

import (
	"log"
	"net/http"
	"strconv"

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

func (h *Handler) Home(ctx *gin.Context) {
	userID := ctx.Param("user")

	log.Println("home request")

	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Print("incorrect id")
	}
	tmpl, home, err := h.app.Home(ctx.Request.Context(), id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(ctx.Writer, *home)
	if err != nil {
		log.Println("unable to render page")
		log.Println(err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) Task(ctx *gin.Context) {
	userID := ctx.Param("user")
	taskID := ctx.Param("task")

	log.Println("task request")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Print("incorrect id")
	}
	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		log.Print("incorrect id")
	}
	tmpl, home, err := h.app.Task(ctx.Request.Context(), userIDInt, taskIDInt)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(ctx.Writer, *home)
	if err != nil {
		log.Println("unable to render page")
		log.Println(err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.Status(http.StatusOK)
}

//func (h *Handler) submit(ctx *gin.Context) {
//}
