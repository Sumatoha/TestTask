package handlers

import (
	"TestTask/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	useCases *service.UseCase
	R        *gin.Engine
}

func NewHandler(useCases *service.UseCase) *handler {
	return &handler{
		useCases: useCases,
		R:        gin.Default(),
	}
}
func (h *handler) Register() {
	h.R.RedirectTrailingSlash = true

	h.R.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "method not allowed",
		})
	})
	h.R.HandleMethodNotAllowed = true
	h.R.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
		})
	})

	user := h.R.Group("/user")
	{
		user.GET(fmt.Sprintf("/redirects"), h.Redirect)
	}

	admin := h.R.Group("/admin")
	{
		admin.GET(fmt.Sprintf("/redirects"), h.GetAllPaths)
		admin.POST(fmt.Sprintf("/redirects"), h.AddPath)
		admin.GET(fmt.Sprintf("/redirects/:%s", addressPath), h.GetPath)
		admin.PATCH(fmt.Sprintf("/redirects/:%s", addressPath), h.UpdatePath)
		admin.DELETE(fmt.Sprintf("/redirects/:%s", addressPath), h.DeletePath)
	}
}
