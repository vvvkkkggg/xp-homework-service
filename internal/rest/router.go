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

	r.GET("/:user/home", h.home)
	r.GET("/:user/tasks/:task", h.task)
	r.POST("/:user/tasks/:task", h.postTask)

	/*
		name := c.Param("name")
				c.String(http.StatusOK, "Hello %s", name)
	*/

	return r
}
