package handlers

import (
	"delivery/constants"
	"delivery/entities"
	"delivery/pkg/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Test(c *gin.Context) {
	h.handleResponse(c, http.OK, "ishladi")
}

func (h *Handler) CreateXozmak(c *gin.Context) {
	var body entities.Xozmak
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	body.ID = uuid.NewString()
	body.CreatedBy = entities.NullString("ab89ca99-3c18-4751-9c07-51a2ee85751e")

	err = h.adminController.CreateXozmak(c, body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, constants.Success)
}
