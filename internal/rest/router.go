package rest

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(h *Handler) *Router {
	gin.SetMode(gin.ReleaseMode)
	r := &Router{
		Engine: gin.New(),
	}

	r.GET("/:user/home", h.Home)
	r.GET("/:user/tasks/:task", h.Task)
	//r.POST("/:user/tasks/:task", h.submit)	todo: submit

	return r
}
