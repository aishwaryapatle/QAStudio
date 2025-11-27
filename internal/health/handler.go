package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "Welcome to Testora backend (minimal)",
		"status":  "ready",
	})
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}