package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hvmidrezv/web-app/api/helper"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Server is Working", true, 0))
	return
}
