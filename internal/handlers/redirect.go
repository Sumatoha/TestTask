package handlers

import (
	"TestTask/internal/handlers/dto"
	"TestTask/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

const addressPath = "address"

func (h *handler) GetAllPaths(c *gin.Context) {
	res, err := h.useCases.PathService.GetAllPaths(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (h *handler) GetPath(c *gin.Context) {
	address := c.Param(addressPath)

	res, err := h.useCases.PathService.GetPathByID(c.Request.Context(), address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (h *handler) DeletePath(c *gin.Context) {
	address := c.Param(addressPath)

	err := h.useCases.PathService.DeletePath(c.Request.Context(), address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.Status(http.StatusOK)
}

func (h *handler) UpdatePath(c *gin.Context) {
	address := c.Param(addressPath)
	req := dto.UpdatePathRequest{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	params := models.Path{
		ActiveLink:  req.ActiveLink,
		HistoryLink: req.HistoryLink,
	}

	err := h.useCases.PathService.UpdatePath(c.Request.Context(), address, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.Status(http.StatusOK)
}

func (h *handler) AddPath(c *gin.Context) {
	req := dto.AddPathRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	params := models.Path{
		ActiveLink:  req.ActiveLink,
		HistoryLink: req.HistoryLink,
	}

	id, err := h.useCases.PathService.AddPath(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, id)
}

func (h *handler) Redirect(c *gin.Context) {
	link, ok := c.GetQuery("link")
	if !ok {
		c.JSON(http.StatusBadRequest, "invalid link")
	}
	newLink, valid := h.useCases.PathService.RedirectPath(c.Request.Context(), link)
	if !valid {
		c.JSON(http.StatusMovedPermanently, newLink)
		return
	}
	c.Status(http.StatusOK)
}
