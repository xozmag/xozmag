package handlers

import (
	"delivery/constants"
	"delivery/entities"
	"delivery/pkg/http"
	"delivery/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateXozmak(c *gin.Context) {
	var body entities.Xozmak
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, constants.BadRequest)
		return
	}
	body.ID = uuid.NewString()
	body.CreatedBy = entities.NullString("ab89ca99-3c18-4751-9c07-51a2ee85751e")

	err = h.adminController.CreateXozmak(c, body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}

	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) GetXozmak(c *gin.Context) {
	data, err := h.adminController.GetXozmak(c)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
	}
	h.handleResponse(c, http.OK, data)
}

func (h *Handler) UpdateXozmak(c *gin.Context) {
	var body entities.Xozmak
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, constants.BadRequest)
		return
	}
	body.ID = c.Param("id")

	if !utils.IsValidUUID(body.ID) {
		h.handleResponse(c, http.BadRequest, "Invalid UUID format") 
		return
	}

	err = h.adminController.UpdateXozmak(c, body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
	}
	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) DeleteXozmak(c *gin.Context) {
	id := c.Param("id")

	if !utils.IsValidUUID(id) {
		h.handleResponse(c, http.BadRequest, "Invalid UUID format") 
		return
	}

	err := h.adminController.DeleteXozmak(c, id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
	}
	h.handleResponse(c, http.OK, constants.Success)
}
