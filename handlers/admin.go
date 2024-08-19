package handlers

import (
	"delivery/pkg/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) Test(c *gin.Context) {
	h.handleResponse(c, http.OK, "ishladi")
}